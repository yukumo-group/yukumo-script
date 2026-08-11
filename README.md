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
| CLI | `just build-cli` | `dist/yukumo.exe` (Windows) / `dist/yukumo` |
| clib | `just build-clib` | `dist/clib/yukumo.dll` + `yukumo.h` (Windows); `.so` / `.dylib` on Unix |
| both | `just build` | CLI + clib |

```bash
just build        # clib then CLI
just build-cli    # standalone executable only
just build-clib   # shared library for other programs
just clean        # remove dist/
```

Or with plain Go:

```bash
go build -o dist/yukumo.exe ./cmd/yukumo
go build -buildmode=c-shared -o dist/clib/yukumo.dll ./cmd/clib
```

### Prerequisites (Windows)

Place AquesTalk2 `.dll` / `.lib` under `third_party/aquestalk2/win64/`. Audio generation (CLI and clib) requires Windows + AquesTalk2. The native DLL is loaded at runtime from `third_party/aquestalk2/win64/AquesTalk2.dll` (working directory should be the repo root).

### clib C API

Header is generated next to the library (`dist/clib/yukumo.h`). Exports include:

- `YukumoInit` — runtime dirs, phont map, characters, tasks
- `YukumoListPhonts` / `YukumoListTasks` — name lists (`StringList`)
- `YukumoGenerateByPhont` — generate wav by phont (Windows)
- `YukumoFreeString` / `YukumoFreeStringList` / `YukumoFreeErrorMessage` — free C memory

Runtime I/O lives under `runtime/` (`data`, `phonts`, `wav`, `result`, `examples`).
