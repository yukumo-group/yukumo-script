# yukumo-script

A Go application that generates Yukumo audio from script using AquesTalk2.

## Layout

```text
cmd/yukumo/              CLI entrypoint
internal/cmdinterface/   CLI-only helpers
pkg/                     Shared libraries (also used by yukumo-script-gui)
third_party/aquestalk2/  Native AquesTalk2 libraries (local)
runtime/                 Runtime I/O (phonts, wav, result, datas, examples)
yukumo-script-gui/       Wails GUI submodule
```

## Build (Windows)

Place AquesTalk2 `.dll` / `.lib` under `third_party/aquestalk2/win64/`, then:

```bash
go build -o yukumo.exe ./cmd/yukumo
```

Runtime directories are created under `runtime/` on first run.
