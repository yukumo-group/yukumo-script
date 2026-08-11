# yukumo-script

A Go application that generates Yukumo audio from script using AquesTalk2.

## Layout

```text
cmd/yukumo/              CLI entrypoint (standalone)
cmd/clib/                C shared library exports (-buildmode=c-shared)
pkg/api/                 Shared helpers used by CLI and clib
pkg/                     Shared libraries (also used by yukumo-script-gui)
internal/cmdinterface/   CLI-only helpers
third_party/aquestalk2/  Native AquesTalk2 libraries (local)
runtime/                 Runtime I/O (phonts, wav, result, data, examples)
dist/                    Build outputs (gitignored)
yukumo-script-gui/       Wails GUI submodule
```

## Build targets

Two independent artifacts (CLI does **not** load the shared library):

| Target | Recipe | Output |
|--------|--------|--------|
| CLI | `just build-cli` | `dist/cli/{win64,win32,linux64,linux32,macos}/yukumo[.exe]` |
| clib | `just build-clib` | `dist/clib/{win64,win32}/yukumo.dll`, `…/linux*/libyukumo.so`, `…/macos/libyukumo.dylib` + `yukumo.h` |
| both | `just build` | CLI + clib |

```bash
just build              # all platform CLIs + clibs
just build-cli          # win64/win32 + linux64/linux32 + macos CLI
just build-clib         # matching shared libraries
just build-cli-linux64  # single platform (also -win32, -macos, …)
just clean              # remove dist/ and build/
```

Cross linux/macOS builds use [Zig](https://ziglang.org/) (`ZIG=zig`). Windows 32-bit needs MinGW (`CC32=i686-w64-mingw32-gcc`). Linux/macOS CLI crosses use `-tags noaudio` (playback stubbed; generation still works).

Or with plain Go (host only):

```bash
go build -o dist/cli/win64/yukumo.exe ./cmd/yukumo
go build -tags clib -buildmode=c-shared -o dist/clib/win64/yukumo.dll ./cmd/clib
```

### Prerequisites (AquesTalk2)

Place the AquesTalk2 SDK under `third_party/aquestalk2/{win,linux,mac}/` (vendor layout with `lib` / `lib64` and `phont`). Audio generation loads the matching shared library at runtime from the working directory (repo root):

- Windows amd64: `third_party/aquestalk2/win/lib64/AquesTalk2.dll`
- Windows 386: `third_party/aquestalk2/win/lib/AquesTalk2.dll`
- Linux amd64: `third_party/aquestalk2/linux/lib64/libAquesTalk2Eva.so.2.3`
- Linux 386: `third_party/aquestalk2/linux/lib/libAquesTalk2Eva.so.2.3`
- macOS: `third_party/aquestalk2/mac/lib/libAquesTalk2Eva.dylib`

### clib C API

Header is generated next to the library (`dist/clib/yukumo.h`). Exports include:

- `YukumoInit` — runtime dirs, phont map, characters, tasks
- `YukumoListPhonts` / `YukumoListTasks` — name lists (`StringList`)
- `YukumoGenerateByPhont` — generate wav by phont
- `YukumoFreeString` / `YukumoFreeStringList` / `YukumoFreeErrorMessage` — free C memory

Runtime I/O lives under `runtime/` (`data`, `phonts`, `wav`, `result`, `examples`).
