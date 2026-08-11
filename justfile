# Dual build targets: standalone CLI + C shared library (clib).

set windows-shell := ["powershell.exe", "-NoProfile", "-Command"]

os_family := os()

exe_ext := if os_family == "windows" { ".exe" } else { "" }

clib_name := if os_family == "windows" {
    "yukumo.dll"
} else if os_family == "macos" {
    "libyukumo.dylib"
} else {
    "libyukumo.so"
}

clib_out := "dist/clib/" + clib_name
cli_out := "dist/yukumo" + exe_ext

default: build

# Build both clib and CLI (order is convenience only; CLI does not link clib).
build: build-clib build-cli

# Build C shared library + generated header under dist/clib/.
[windows]
build-clib:
    New-Item -ItemType Directory -Force -Path dist/clib | Out-Null
    go build -buildmode=c-shared -o {{clib_out}} ./cmd/clib

[unix]
build-clib:
    mkdir -p dist/clib
    go build -buildmode=c-shared -o {{clib_out}} ./cmd/clib

# Build standalone CLI executable under dist/.
[windows]
build-cli:
    New-Item -ItemType Directory -Force -Path dist | Out-Null
    go build -o {{cli_out}} ./cmd/yukumo

[unix]
build-cli:
    mkdir -p dist
    go build -o {{cli_out}} ./cmd/yukumo

# Remove build outputs.
[windows]
clean:
    if (Test-Path dist) { Remove-Item -Recurse -Force dist }

[unix]
clean:
    rm -rf dist
