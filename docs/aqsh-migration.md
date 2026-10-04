# 合成层切换到 aqtk-shim（Aqsh）—— 变更对照

> 状态：**设计待评审**（未动代码）。约束：**Go 侧零 C** —— 不写 `.c`/`.h` 桥接，不 `#include`，不用 cgo。
> 目标：合成改为调用 aqtk-shim 的 `AqShim_Synthesize`，运行期从 `third-party/<platform>/` 加载；**用单函数取代 `NewGenerator` 对象**。

## 1. 现状（为什么非改不可）

| 事实 | 位置 |
|---|---|
| 自带 C 桥接，只绑 AquesTalk2 单引擎 | [`YukumoGenerator.c:140-181`](../internal/generator/aquestalk2/YukumoGenerator.c#L140-L181) |
| `libPath()` 指向 `third_party/aquestalk2/…`，该目录**不存在** → 当前任何平台都合成不出音频 | [`libpath.go:9-33`](../internal/generator/aquestalk2/libpath.go#L9-L33) |
| `Aqsh` 已 stage 成 `third-party/{win64,win32,linux64,linux32}/` + `Aqsh.h` | [`install.cmake:10-16`](../aqtk-shim/cmake/install.cmake#L10-L16) |
| 引擎库按 **Aqsh 自身所在目录** 解析，与 CWD 无关 | Win `/DEPENDENTLOADFLAG:0x900`（[`windows.cmake:93`](../aqtk-shim/cmake/platform/windows.cmake#L93)）；Linux `$ORIGIN`；macOS `@loader_path` |

调用面只有 3 处，形状统一：`aquestalk2.NewGenerator(speed, phontPath, resultPath, text).GenerateWav()`（[`singlesentence_task.go:221`](../internal/generator/tasks/singlesentence/singlesentence_task.go#L221)、[`chorus.go:114`](../internal/generator/tasks/chorus/chorus.go#L114)、[`script.go:59`](../internal/example/script.go#L59)）。

**为什么要连带抛弃这个对象形态**：`internal/generator/generator.go` 的单方法接口 `Generator{ GenerateWav() error }` 在**全仓零引用**（没有任何文件 import `internal/generator`，也没有变量/字段/返回值用过它），`NewGenerator` 返回的还是具体类型而非接口，实现只有一个。也就是说：构造函数把参数从方法签名搬进字段，只是在为一个没人用的多态点服务。而 Aqsh 已经在 C 层用 `AqVersion version` 一个判别式字段统一了三引擎——判别式字段取代类型层级，Go 侧再搞每引擎一个类型就没有意义了。
附带代价：`resultPath` 被塞进构造参数，把「合成」和「落盘」耦成一个类型（换成内存缓冲就必须改 API）；`Speed`/`PhontName` 在 `Task` 与临时对象里各存一份会漂移；对象不持有任何句柄（C 侧是进程级 static），并发时反而给人「每实例独立」的错觉。

## 2. 变更对照表

| 位置 | 原来干什么 | 怎么改 | 之后干什么 |
|---|---|---|---|
| `aquestalk2/YukumoGenerator.c` / `.h` | 手写 C：`dlopen` AquesTalk2、解析 2 个符号、`fopen` 写 wav | **删** | 桥接整体消失 |
| `aquestalk2/aquestalk2.go` | cgo 绑定 + `Generator` 对象，硬编码 `-1..-4` 错误 | **删** | 由 `aqsh.Synthesize` 单函数取代 |
| `aquestalk2/libpath.go` | 返回 `third_party/aquestalk2/…`（已失效） | **删**，逻辑移入 `aqsh/locate.go` | `third-party/<arch>/` 多候选搜索 |
| `internal/generator/generator.go` | 单方法接口（零引用死亡抽象） | **删** | 不再需要多态点 |
| **新** `aqsh/abi.go` | — | 新增 | `AqSynthParam` 的 Go 字段镜像 + 枚举；零 C 下的编译期契约 |
| **新** `aqsh/load_unix.go` | — | 新增 | `purego.Dlopen` + `RegisterFunc` 绑定 5 个符号 |
| **新** `aqsh/load_windows.go` | — | 新增 | `LoadLibraryEx` + `GetProcAddress` + `SyscallN` |
| **新** `aqsh/locate.go` + `aqsh.go` | — | 新增 | 定位模块；`Load` / `Params` / `Synthesize` / `Shutdown` |
| 3 个调用点 | `NewGenerator(...).GenerateWav()` | 换成 `aqsh.Synthesize(...)` + 自己 `os.WriteFile` | 每个调用点约 3 行，`Task.Generate` 砍掉一半 |
| `pkg/api/api.go:24` | 示例音频生成失败 → `Init()` 整体失败 | 降级为非致命；新增 `SetAqshDir(path)` 供 clib 宿主 | 缺库不再炸启动 |
| `pkg/utils/consts.go` | — | 加 `ThirdPartyDir = "third-party"` | 统一路径口径 |
| `go.mod` | `purego`、`x/sys` 是 indirect | 提为 direct | 直接依赖 |
| `justfile` / `just/cli.just` / `just/clib.just` | build 只产二进制 | 加 `stage-cli-<p>`：把 `third-party/<p>/*` 拷到 `dist/.../<p>/`，并让 build 依赖它 | 发布形态 = exe + Aqsh + 引擎库同目录 |
| `README.md:38,47-55` | 描述 `third_party/aquestalk2/` 旧布局；称 macOS 交叉"生成仍可用" | 重写为 `just init <platform>` + `third-party/<platform>/`；删掉 macOS 那句错误表述 | 文档与实现一致 |
| CI `lint.yaml` / `coverage.yaml` | checkout 不拉子模块 | 可选 `submodules: true` | 让 ABI 测试能读到 `aqtk-shim/src/Aqsh.h` |
| **新** `tests/aqsh/{locate,abi,synthesize_smoke}_test.go` | — | 新增 | 定位表 / ABI 漂移 / 真合成（缺库则 skip） |

顺带消失：`#cgo windows LDFLAGS: -lkernel32`、`#cgo linux LDFLAGS: -ldl`。

**顺带发现的既有问题**（与合成无关，但建议先确认）：

1. [`cmd/go.mod:38`](../cmd/go.mod#L38) 的 `replace github.com/yukumo/yukumo-script => ../` 与 require 的 `github.com/yukumo-group/yukumo-script` 路径不一致 → replace 不生效，且无 `go.work` 兜底 → CLI 解析的是已发布伪版本，**不是工作树**。
2. justfile 构建 `./cmd/yukumo`（[`cli.just:7`](../just/cli.just#L7)），但该目录不存在（`cmd/` 自身就是独立模块的 `package main`）。
3. **`chorus.Task.Generate` 会 panic**：它在两个 worker 还没发完时就 `close(resultChan)`（[`chorus.go:318`](../internal/generator/tasks/chorus/chorus.go#L318)），worker 随后发送（[:297](../internal/generator/tasks/chorus/chorus.go#L297)、[:313](../internal/generator/tasks/chorus/chorus.go#L313)）→ `send on closed channel`（goroutine 里的 panic 无法 recover，直接崩进程）；且 `group.Wait()` 从未调用，worker 错误被吞、`MixAudios` 可能拿到不完整列表。触发路径：sequence 中含 chorus 句子 → [`sentence.go:212`](../internal/generator/tasks/sequence/sentence.go#L212) 产出 `*chorus.Task` → [`sequence.go:220`](../internal/generator/tasks/sequence/sequence.go#L220) 经 `tasks.Task` 接口调用。修法照抄 [`sequence.go:229`](../internal/generator/tasks/sequence/sequence.go#L229)：先 `group.Wait()` 再 `close`（缓冲容量恰为最大发送数，不会死锁）。

## 3. 并发：现状，以及合成层必须假定什么

**上游怎么并发**：

| 位置 | 结构 | 并发合成数 | 正确性 |
|---|---|---|---|
| [`example/script.go:23-24`](../internal/example/script.go#L23-L24) | `errgroup` + `SetLimit(NumCPU*2)`，末尾 `Wait()` | ≤ NumCPU×2 | ✅（结果写 `syncutils.Map`，有锁） |
| [`sequence.go:209-229`](../internal/generator/tasks/sequence/sequence.go#L209-L229) | `errgroup` + `SetLimit(NumCPU*2)`，每句子一个 goroutine | ≤ NumCPU×2 | ✅ `Wait()` 后再 combine |
| [`chorus.go:141,195`](../internal/generator/tasks/chorus/chorus.go#L141) | 外层 2 个 goroutine ＋ 内层各 `SetLimit(NumCPU*2)` | 最坏 ~NumCPU×4 | ❌ 见 §2 问题 3 |
| [`singlesentence_task.go:159`](../internal/generator/tasks/singlesentence/singlesentence_task.go#L159) | 无并发，顺序合成 | 1 | ✅ |
| 输出文件名 | `uuid` + `UnixNano`（[`singlesentence_task.go:101`](../internal/generator/tasks/singlesentence/singlesentence_task.go#L101)、[`chorus.go:74`](../internal/generator/tasks/chorus/chorus.go#L74)） | — | ✅ 并发不撞名 |

结论：**合成层在正常路径上会被最多 ~NumCPU×4 个 goroutine 同时调用**（sequence→chorus 嵌套），启动时 `api.Init()` 的示例生成还会再加 NumCPU×2 路。所以必须假定"高并发调用"。

**单函数 vs 对象，与并发的关系**：纯函数没有共享的可变接收者、没有半初始化对象、没有"谁来 Close"，从 Go 语义看是并发的**更好**形态；对象形态只会制造"每实例独立"的错觉。但**函数签名不提供任何并发安全**——Aqsh 背后是进程级可变全局状态（引擎库内部静态、`wav_track` 指针表、`aqtk1` 的 voice registry），`Synthesize` 是**外观纯、实现不纯**。安全只能来自锁或每 goroutine 独立引擎实例。

**重构时的注意清单**：

1. **决定锁放哪**（最重要）。三选一：① Go 侧全局 `synthMu` —— 最省事，但把上游的并行全部串行化，合成是 CPU 密集，吞吐降到 1/N；② Go 侧**分桶锁**（按 `Engine`，AQ1 再按 `preset`，因为不同 voice 是不同 DLL）—— 保住大部分并行度且不动子模块，**推荐 P3 用这个**；③ 修 shim 在 `aqtk1/2/10` 各自加锁（一级锁）—— 最优，但要动子模块并出一次 release，作为后续 issue。
2. **`Load` 用 `sync.Once`，不要在并发路径重试**。定位失败要**缓存失败**（或只暴露显式 `Reset()`）；否则用户后补 `just init` 时会在并发下反复 dlopen/dlclose（Linux 上 dlclose 可能卸载仍在使用的库）。
3. **`Shutdown` 绝不能在并发合成中途调用**：它释放 Aqsh 侧**所有**已 track 的 WAV（[`wav_track.cpp:41-49`](../aqtk-shim/src/wav_track.cpp#L41-L49)）。建议只在进程退出时调，或干脆不调（每次都已 `FreeWav`）。
4. **拷完立刻 `FreeWav` 并置 nil**。`wav_track` 表有锁（[`wav_track.cpp:8`](../aqtk-shim/src/wav_track.cpp#L8)），Free 与其它合成可并发；但同一指针重复 Free 是**静默 return**（[`wav_track.cpp:30-33`](../aqtk-shim/src/wav_track.cpp#L30-L33)），错了不会报错，只能在 Go 侧保证一次性。
5. **`ctx` 只能用来"别排队"，不能中断 C 调用**。合成是不可中断的 C 调用；建议 `Synthesize(ctx, …)` 在**拿锁前**检查 `ctx.Err()`，让取消后的 goroutine 不要白排队。
6. **同一个 `Task` 防重入**：现在并发调两次 `Generate()` 会跑两遍合成（`task.Lock()` 只保护 `ResultFile` 赋值）。重构时顺手加"已生成则跳过"或 per-task 生成锁。
7. **错误码取文本可并发**：`AqShim_ErrorMessage` 内部有锁（[`aq_error.cpp:74`](../aqtk-shim/src/aq_error.cpp#L74)），但别把返回的 `const char*` 留成裸指针跨调用。
8. **P4 混合引擎后锁要分桶**：同一进程出现 AQ1/AQ2/AQ10 混跑时，全局锁也不会错，只是更慢；分桶按 `Engine`+`preset` 才真正有效。

## 4. 外置 API 文档

| 用途 | 链接 |
|---|---|
| Aqsh C 接口（权威定义、参数语义） | [`aqtk-shim/src/Aqsh.h`](../aqtk-shim/src/Aqsh.h#L75-L115) · [GitHub](https://github.com/yukumo-group/aqtk-shim/blob/main/src/Aqsh.h) |
| 参数与错误码说明 | [`README.zh.md`](../aqtk-shim/README.zh.md#L98-L152) |
| purego：`Dlopen`/`Dlsym`/`RegisterFunc`/`SyscallN` | [pkg.go.dev/github.com/ebitengine/purego](https://pkg.go.dev/github.com/ebitengine/purego) |
| purego：类型映射 / 内存生命周期 / 结构体限制 | [Type Conversions](https://pkg.go.dev/github.com/ebitengine/purego#hdr-Type_Conversions__Go_____C_-RegisterFunc) · [Memory](https://pkg.go.dev/github.com/ebitengine/purego#hdr-Memory) · [Structs](https://pkg.go.dev/github.com/ebitengine/purego#hdr-Structs) |
| purego 为何不能用于 Windows | [v0.9.0 `dlfcn.go` 的 build tag](https://github.com/ebitengine/purego/blob/v0.9.0/dlfcn.go)（`darwin \|\| freebsd \|\| linux \|\| netbsd`） |
| Windows 侧：`LoadLibraryEx` / `GetProcAddress` / `LOAD_WITH_ALTERED_SEARCH_PATH` | [pkg.go.dev/golang.org/x/sys/windows](https://pkg.go.dev/golang.org/x/sys/windows) |
| 调用任意函数指针 | [syscall.SyscallN](https://pkg.go.dev/syscall#SyscallN) |
| 布局校验 | [unsafe.Offsetof / Sizeof](https://pkg.go.dev/unsafe#Offsetof) |

## 5. 核心代码（省略号部分为重复样板）

`abi.go` —— **不要手写 padding**：C 的 `int` 是 4 字节、指针按指针宽度对齐，Go 的 `int32`/`unsafe.Pointer` 规则相同，因此 64 位得 80 字节、32 位得 52 字节，编译器自己算对。

```go
type synthParam struct {
	version, preset, speed int32
	phont, phontData       unsafe.Pointer            // char* / unsigned char*
	base, volume, pitch, accent, lmd, fsc int32      // 仅 AQ10_CUSTOM 生效
	devKey, userKey        unsafe.Pointer            // char*
}
```

`load_unix.go`（`//go:build !windows`）：

```go
handle, err := purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_LOCAL)
addr, err := purego.Dlsym(handle, "AqShim_Synthesize") // 先 Dlsym 探测：RegisterLibFunc 找不到会 panic
purego.RegisterFunc(&s.synthesize, addr)               // func(text string, p unsafe.Pointer, out **byte, n *int32) int32
```

`load_windows.go`（`//go:build windows`）：

```go
h, err := windows.LoadLibraryEx(p, 0, windows.LOAD_WITH_ALTERED_SEARCH_PATH)
addr, err := windows.GetProcAddress(h, "AqShim_Synthesize")
s.synthesize = func(text string, p unsafe.Pointer, out **byte, n *int32) int32 {
	t, _ := syscall.BytePtrFromString(text)
	r, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(t)), uintptr(p),
		uintptr(unsafe.Pointer(out)), uintptr(unsafe.Pointer(n)))
	return int32(r)
}
```

唯一入口与调用点：

```go
// aqsh：核心唯一入口（锁按 §3.1 的分桶方案实现）
func Synthesize(ctx context.Context, p Params, text string) ([]byte, error)

// 调用点（原 NewGenerator(...).GenerateWav() 的位置）
wav, err := aqsh.Synthesize(ctx, aqsh.Params{Engine: aqsh.EngineAq2, Speed: int32(speed), PhontPath: phontPath}, processedText)
if err != nil { return err }
if err := os.WriteFile(fileName, wav, 0o644); err != nil { return err }
```

> 关键有利条件：这 5 个函数签名里**没有按值传递的结构体**，全是指针/整数，正好避开 [purego 的 Structs 限制](https://pkg.go.dev/github.com/ebitengine/purego#hdr-Structs)。

## 6. 必须知道的坑

1. **结构体不要手写 padding** —— 写了会在 386 上错位；靠 `unsafe.Sizeof` 的 80/52 做哨兵断言。
2. **释放只能用 `AqShim_FreeWav`** —— Aqsh 按指针查表选引擎自己的 free（[`wav_track.cpp:23-39`](../aqtk-shim/src/wav_track.cpp#L23-L39)），Go 侧不能 free、不能交给 GC、不能重复释放。
3. **并发安全要靠锁，不能靠函数式签名** —— 见 §3.1 的三选一；shim 只锁了 `wav_track`/`aq_error`，引擎调用无锁（[`aqtk2.cpp:88`](../aqtk-shim/src/aqtk2.cpp#L88)、[`aqtk10.cpp:84`](../aqtk-shim/src/aqtk10.cpp#L84)、[`aqtk1.cpp:103-118`](../aqtk-shim/src/aqtk1.cpp#L103-L118)）。
4. **ABI 无编译器把关** —— 零 C 意味着字段顺序错了不是编译错误而是静默读错内存，`abi_test.go`（解析头文件算偏移比对 `unsafe.Offsetof`）是必需项，不是可选项。
5. **`phont/` 没 stage** —— `just init` 固定传 `-DAQTK_STAGE_PHONTS=OFF`（[`init-platform.ps1:28`](../just/init-platform.ps1#L28)），`preset` 路径不可用；yukumo 一直传显式音色路径，行为不变。
6. **错误码语义变宽** —— 旧实现压成 `-1..-4`，新实现透出引擎码（负数）+ `AqShim_ErrorMessage` 文本；依赖旧字符串的上层要复核。
7. **macOS 仍不可用** —— `third-party/macos/` 只能在 Mac 上 stage，且 Aqsh 是 arm64-only 而交叉产物是 `darwin/amd64`。
8. **clib 宿主路径不可控** —— `os.Executable()` 在 `c-shared` 里返回宿主 exe 路径，GUI 场景需 `YUKUMO_AQSH_DIR` / `SetAqshDir`。

## 7. 落地顺序

| 阶段 | 内容 | 验收 |
|---|---|---|
| P1 | 新增 `aqsh` 包（`abi.go`/`locate.go`/两个 load 文件/单测），**不动调用点** | `go test -tags "test noaudio" ./tests/...` 绿；`locate_test`、`abi_test` 绿 |
| P2 | 加真合成 smoke 测试（含序列化/分桶锁的并发用例，`-race`） | 本机 win64 下 WAV 头/尺寸断言通过；`-race` 无报告；缺库 skip |
| P3 | 3 个调用点改用 `Synthesize`，删 `aquestalk2/` **与 `generator.go` 接口**；改 README 与 `Init()` 降级 | 全平台构建通过；CLI 端到端出一条单句（需先解决 §2 的 1、2） |
| P4 | 扩参数面：`Character`/`Task` 加 `engine`/`preset` + aq10 六参数（JSON 缺省 = AQ2），CLI 暴露；锁改分桶 | 三引擎各出一条（含同进程混跑）；旧 `characters.json` 仍可加载 |

**待拍板**：① 是否先修 §2 的三处既有问题（决定 P3 能否端到端验收）② 锁方案选 §3.1 的 ①/②/③ ③ P4 参数落点（角色存引擎，还是独立「音色」概念）。
