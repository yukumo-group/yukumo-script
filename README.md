# yukumo-script

A Go application that generates Yukumo audio from script using AquesTalk2.

## Layout

```text
cmd/yukumo/              CLI entrypoint
internal/cmdinterface/   CLI-only helpers
pkg/                     Shared libraries (also used by yukumo-script-gui)
third_party/aquestalk2/  Native AquesTalk2 libraries (local)
runtime/                 Runtime I/O (phonts, wav, result, data, examples)
yukumo-script-gui/       Wails GUI submodule
```

## Build (Windows)

Place AquesTalk2 `.dll` / `.lib` under `third_party/aquestalk2/win64/`, then:

```bash
go build -o yukumo.exe ./cmd/yukumo
go run ./cmd/yukumo
```

The native DLL is loaded at runtime from `third_party/aquestalk2/win64/AquesTalk2.dll` (working directory should be the repo root). Runtime I/O lives under `runtime/` (`data`, `phonts`, `wav`, `result`, `examples`).
