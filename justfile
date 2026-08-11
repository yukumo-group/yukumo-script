# Build CLI + clib for win64/win32, linux64/linux32, and macos (Zig cross).

set windows-shell := ["powershell.exe", "-NoProfile", "-Command"]

cc_host := env_var_or_default("CC_HOST", "gcc")
cc32 := env_var_or_default("CC32", "i686-w64-mingw32-gcc")
zig := env_var_or_default("ZIG", "zig")
cc_linux64 := env_var_or_default("CC_LINUX64", zig + " cc -target x86_64-linux-gnu")
cc_linux32 := env_var_or_default("CC_LINUX32", zig + " cc -target x86-linux-gnu")
cc_macos := env_var_or_default("CC_MACOS", zig + " cc -target x86_64-macos-none")
cc_win64_cross := zig + " cc -target x86_64-windows-gnu"
cc_win32_cross := zig + " cc -target x86-windows-gnu"

# -buildid= skips Darwin UUID rewrite (broken on Windows hosts).
ld_release := "-s -w -buildid="
ld_debug := "-buildid="
gc_debug := "all=-N -l"
cflags_debug := "-g -O0"

clib_tags := "clib"
cli_cross_tags := "noaudio"
macos_stub_dir := "build/macos-stubs"
macos_ldflags := "-L" + macos_stub_dir + " -Wl,-undefined,dynamic_lookup"

default: build

build: build-clib build-cli
build-debug: build-clib-debug build-cli-debug

build-clib: build-clib-win64 build-clib-win32 build-clib-linux64 build-clib-linux32 build-clib-macos
build-cli: build-cli-win64 build-cli-win32 build-cli-linux64 build-cli-linux32 build-cli-macos
build-clib-debug: build-clib-win64-debug build-clib-win32-debug build-clib-linux64-debug build-clib-linux32-debug build-clib-macos-debug
build-cli-debug: build-cli-win64-debug build-cli-win32-debug build-cli-linux64-debug build-cli-linux32-debug build-cli-macos-debug

# ---- release ----

[windows]
build-clib-win64: (_clib-win "windows" "amd64" cc_host "dist/clib/win64/yukumo.dll" ld_release "" "")
[unix]
build-clib-win64: (_clib-unix "windows" "amd64" cc_win64_cross "dist/clib/win64/yukumo.dll" ld_release "" "")

[windows]
build-clib-win32: (_clib-win "windows" "386" cc32 "dist/clib/win32/yukumo.dll" ld_release "" "")
[unix]
build-clib-win32: (_clib-unix "windows" "386" cc_win32_cross "dist/clib/win32/yukumo.dll" ld_release "" "")

[windows]
build-clib-linux64: (_clib-win "linux" "amd64" cc_linux64 "dist/clib/linux64/libyukumo.so" ld_release "" "")
[unix]
build-clib-linux64: (_clib-unix "linux" "amd64" cc_linux64 "dist/clib/linux64/libyukumo.so" ld_release "" "")

[windows]
build-clib-linux32: (_clib-win "linux" "386" cc_linux32 "dist/clib/linux32/libyukumo.so" ld_release "" "")
[unix]
build-clib-linux32: (_clib-unix "linux" "386" cc_linux32 "dist/clib/linux32/libyukumo.so" ld_release "" "")

[windows]
build-clib-macos: macos-stubs (_clib-win "darwin" "amd64" cc_macos "dist/clib/macos/libyukumo.dylib" ld_release "" macos_ldflags)
[unix]
build-clib-macos: macos-stubs (_clib-unix "darwin" "amd64" cc_macos "dist/clib/macos/libyukumo.dylib" ld_release "" macos_ldflags)

[windows]
build-cli-win64: (_cli-win "windows" "amd64" cc_host "dist/cli/win64/yukumo.exe" "" ld_release "" "")
[unix]
build-cli-win64: (_cli-unix "windows" "amd64" cc_win64_cross "dist/cli/win64/yukumo.exe" "" ld_release "" "")

[windows]
build-cli-win32: (_cli-win "windows" "386" cc32 "dist/cli/win32/yukumo.exe" "" ld_release "" "")
[unix]
build-cli-win32: (_cli-unix "windows" "386" cc_win32_cross "dist/cli/win32/yukumo.exe" "" ld_release "" "")

[windows]
build-cli-linux64: (_cli-win "linux" "amd64" cc_linux64 "dist/cli/linux64/yukumo" cli_cross_tags ld_release "" "")
[unix]
build-cli-linux64: (_cli-unix "linux" "amd64" cc_linux64 "dist/cli/linux64/yukumo" cli_cross_tags ld_release "" "")

[windows]
build-cli-linux32: (_cli-win "linux" "386" cc_linux32 "dist/cli/linux32/yukumo" cli_cross_tags ld_release "" "")
[unix]
build-cli-linux32: (_cli-unix "linux" "386" cc_linux32 "dist/cli/linux32/yukumo" cli_cross_tags ld_release "" "")

[windows]
build-cli-macos: macos-stubs (_cli-win "darwin" "amd64" cc_macos "dist/cli/macos/yukumo" cli_cross_tags ld_release "" macos_ldflags)
[unix]
build-cli-macos: macos-stubs (_cli-unix "darwin" "amd64" cc_macos "dist/cli/macos/yukumo" cli_cross_tags ld_release "" macos_ldflags)

# ---- debug ----

[windows]
build-clib-win64-debug: (_clib-win "windows" "amd64" cc_host "dist/debug/clib/win64/yukumo.dll" ld_debug cflags_debug "")
[unix]
build-clib-win64-debug: (_clib-unix "windows" "amd64" cc_win64_cross "dist/debug/clib/win64/yukumo.dll" ld_debug cflags_debug "")

[windows]
build-clib-win32-debug: (_clib-win "windows" "386" cc32 "dist/debug/clib/win32/yukumo.dll" ld_debug cflags_debug "")
[unix]
build-clib-win32-debug: (_clib-unix "windows" "386" cc_win32_cross "dist/debug/clib/win32/yukumo.dll" ld_debug cflags_debug "")

[windows]
build-clib-linux64-debug: (_clib-win "linux" "amd64" cc_linux64 "dist/debug/clib/linux64/libyukumo.so" ld_debug cflags_debug "")
[unix]
build-clib-linux64-debug: (_clib-unix "linux" "amd64" cc_linux64 "dist/debug/clib/linux64/libyukumo.so" ld_debug cflags_debug "")

[windows]
build-clib-linux32-debug: (_clib-win "linux" "386" cc_linux32 "dist/debug/clib/linux32/libyukumo.so" ld_debug cflags_debug "")
[unix]
build-clib-linux32-debug: (_clib-unix "linux" "386" cc_linux32 "dist/debug/clib/linux32/libyukumo.so" ld_debug cflags_debug "")

[windows]
build-clib-macos-debug: macos-stubs (_clib-win "darwin" "amd64" cc_macos "dist/debug/clib/macos/libyukumo.dylib" ld_debug cflags_debug macos_ldflags)
[unix]
build-clib-macos-debug: macos-stubs (_clib-unix "darwin" "amd64" cc_macos "dist/debug/clib/macos/libyukumo.dylib" ld_debug cflags_debug macos_ldflags)

[windows]
build-cli-win64-debug: (_cli-win "windows" "amd64" cc_host "dist/debug/cli/win64/yukumo.exe" "" ld_debug cflags_debug "")
[unix]
build-cli-win64-debug: (_cli-unix "windows" "amd64" cc_win64_cross "dist/debug/cli/win64/yukumo.exe" "" ld_debug cflags_debug "")

[windows]
build-cli-win32-debug: (_cli-win "windows" "386" cc32 "dist/debug/cli/win32/yukumo.exe" "" ld_debug cflags_debug "")
[unix]
build-cli-win32-debug: (_cli-unix "windows" "386" cc_win32_cross "dist/debug/cli/win32/yukumo.exe" "" ld_debug cflags_debug "")

[windows]
build-cli-linux64-debug: (_cli-win "linux" "amd64" cc_linux64 "dist/debug/cli/linux64/yukumo" cli_cross_tags ld_debug cflags_debug "")
[unix]
build-cli-linux64-debug: (_cli-unix "linux" "amd64" cc_linux64 "dist/debug/cli/linux64/yukumo" cli_cross_tags ld_debug cflags_debug "")

[windows]
build-cli-linux32-debug: (_cli-win "linux" "386" cc_linux32 "dist/debug/cli/linux32/yukumo" cli_cross_tags ld_debug cflags_debug "")
[unix]
build-cli-linux32-debug: (_cli-unix "linux" "386" cc_linux32 "dist/debug/cli/linux32/yukumo" cli_cross_tags ld_debug cflags_debug "")

[windows]
build-cli-macos-debug: macos-stubs (_cli-win "darwin" "amd64" cc_macos "dist/debug/cli/macos/yukumo" cli_cross_tags ld_debug cflags_debug macos_ldflags)
[unix]
build-cli-macos-debug: macos-stubs (_cli-unix "darwin" "amd64" cc_macos "dist/debug/cli/macos/yukumo" cli_cross_tags ld_debug cflags_debug macos_ldflags)

# ---- engines (each build must be one line: just runs lines in separate shells) ----

[windows]
_clib-win goos goarch cc out ldflags cflags cgo_ldflags:
    New-Item -ItemType Directory -Force -Path (Split-Path -Parent "{{out}}") | Out-Null
    $ld = '{{cgo_ldflags}}'; if ($ld -like '*-L{{macos_stub_dir}}*') { $ld = $ld.Replace('-L{{macos_stub_dir}}', '-L'+((Resolve-Path '{{macos_stub_dir}}').Path)) }; $env:CGO_ENABLED='1'; $env:GOOS='{{goos}}'; $env:GOARCH='{{goarch}}'; $env:CC='{{cc}}'; if ('{{cflags}}') { $env:CGO_CFLAGS='{{cflags}}' } else { $env:CGO_CFLAGS=$null }; if ($ld) { $env:CGO_LDFLAGS=$ld } else { $env:CGO_LDFLAGS=$null }; if ('{{cflags}}') { go build -tags {{clib_tags}} -buildmode=c-shared -trimpath -ldflags="{{ldflags}}" -gcflags="{{gc_debug}}" -o "{{out}}" ./clib } else { go build -tags {{clib_tags}} -buildmode=c-shared -trimpath -ldflags="{{ldflags}}" -o "{{out}}" ./clib }; $dir = Split-Path -Parent "{{out}}"; $base = [IO.Path]::GetFileNameWithoutExtension("{{out}}"); if (($base -ne 'yukumo') -and (Test-Path "$dir/$base.h")) { Move-Item -Force "$dir/$base.h" "$dir/yukumo.h" }

[unix]
_clib-unix goos goarch cc out ldflags cflags cgo_ldflags:
    mkdir -p "$(dirname "{{out}}")"
    ld="{{cgo_ldflags}}"; case "$ld" in *-L{{macos_stub_dir}}*) ld="$(echo "$ld" | sed "s|-L{{macos_stub_dir}}|-L$(cd {{macos_stub_dir}} && pwd)|")" ;; esac; if [ -n "{{cflags}}" ]; then CGO_ENABLED=1 GOOS={{goos}} GOARCH={{goarch}} CC="{{cc}}" CGO_CFLAGS="{{cflags}}" CGO_LDFLAGS="$ld" go build -tags {{clib_tags}} -buildmode=c-shared -trimpath -ldflags="{{ldflags}}" -gcflags="{{gc_debug}}" -o "{{out}}" ./clib; else CGO_ENABLED=1 GOOS={{goos}} GOARCH={{goarch}} CC="{{cc}}" CGO_LDFLAGS="$ld" go build -tags {{clib_tags}} -buildmode=c-shared -trimpath -ldflags="{{ldflags}}" -o "{{out}}" ./clib; fi; dir="$(dirname "{{out}}")"; base="$(basename "{{out}}" | sed 's/\.[^.]*$//')"; if [ "$base" != "yukumo" ] && [ -f "$dir/$base.h" ]; then mv -f "$dir/$base.h" "$dir/yukumo.h"; fi

[windows]
_cli-win goos goarch cc out tags ldflags cflags cgo_ldflags:
    New-Item -ItemType Directory -Force -Path (Split-Path -Parent "{{out}}") | Out-Null
    $ld = '{{cgo_ldflags}}'; if ($ld -like '*-L{{macos_stub_dir}}*') { $ld = $ld.Replace('-L{{macos_stub_dir}}', '-L'+((Resolve-Path '{{macos_stub_dir}}').Path)) }; $env:CGO_ENABLED='1'; $env:GOOS='{{goos}}'; $env:GOARCH='{{goarch}}'; $env:CC='{{cc}}'; if ('{{cflags}}') { $env:CGO_CFLAGS='{{cflags}}' } else { $env:CGO_CFLAGS=$null }; if ($ld) { $env:CGO_LDFLAGS=$ld } else { $env:CGO_LDFLAGS=$null }; if ('{{tags}}' -and '{{cflags}}') { go build -tags {{tags}} -trimpath -ldflags="{{ldflags}}" -gcflags="{{gc_debug}}" -o "{{out}}" ./cmd/yukumo } elseif ('{{tags}}') { go build -tags {{tags}} -trimpath -ldflags="{{ldflags}}" -o "{{out}}" ./cmd/yukumo } elseif ('{{cflags}}') { go build -trimpath -ldflags="{{ldflags}}" -gcflags="{{gc_debug}}" -o "{{out}}" ./cmd/yukumo } else { go build -trimpath -ldflags="{{ldflags}}" -o "{{out}}" ./cmd/yukumo }

[unix]
_cli-unix goos goarch cc out tags ldflags cflags cgo_ldflags:
    mkdir -p "$(dirname "{{out}}")"
    ld="{{cgo_ldflags}}"; case "$ld" in *-L{{macos_stub_dir}}*) ld="$(echo "$ld" | sed "s|-L{{macos_stub_dir}}|-L$(cd {{macos_stub_dir}} && pwd)|")" ;; esac; tags_flag=""; [ -n "{{tags}}" ] && tags_flag="-tags {{tags}}"; if [ -n "{{cflags}}" ]; then CGO_ENABLED=1 GOOS={{goos}} GOARCH={{goarch}} CC="{{cc}}" CGO_CFLAGS="{{cflags}}" CGO_LDFLAGS="$ld" go build $tags_flag -trimpath -ldflags="{{ldflags}}" -gcflags="{{gc_debug}}" -o "{{out}}" ./cmd/yukumo; else CGO_ENABLED=1 GOOS={{goos}} GOARCH={{goarch}} CC="{{cc}}" CGO_LDFLAGS="$ld" go build $tags_flag -trimpath -ldflags="{{ldflags}}" -o "{{out}}" ./cmd/yukumo; fi

# ---- helpers ----

[windows]
macos-stubs:
    New-Item -ItemType Directory -Force -Path "{{macos_stub_dir}}" | Out-Null
    if (-not (Test-Path "{{macos_stub_dir}}/libresolv.dylib")) { Set-Content "{{macos_stub_dir}}/resolv_stub.c" "void res_9_init(void) {}`n"; & {{zig}} cc -target x86_64-macos-none -shared "{{macos_stub_dir}}/resolv_stub.c" -o "{{macos_stub_dir}}/libresolv.dylib" }
    if (-not (Test-Path "{{macos_stub_dir}}/libpthread.dylib")) { Set-Content "{{macos_stub_dir}}/pthread_stub.c" "void pthread_stub(void) {}`n"; & {{zig}} cc -target x86_64-macos-none -shared "{{macos_stub_dir}}/pthread_stub.c" -o "{{macos_stub_dir}}/libpthread.dylib" }

[unix]
macos-stubs:
    mkdir -p "{{macos_stub_dir}}"
    if [ ! -f "{{macos_stub_dir}}/libresolv.dylib" ]; then printf 'void res_9_init(void) {}\n' > "{{macos_stub_dir}}/resolv_stub.c"; {{zig}} cc -target x86_64-macos-none -shared "{{macos_stub_dir}}/resolv_stub.c" -o "{{macos_stub_dir}}/libresolv.dylib"; fi
    if [ ! -f "{{macos_stub_dir}}/libpthread.dylib" ]; then printf 'void pthread_stub(void) {}\n' > "{{macos_stub_dir}}/pthread_stub.c"; {{zig}} cc -target x86_64-macos-none -shared "{{macos_stub_dir}}/pthread_stub.c" -o "{{macos_stub_dir}}/libpthread.dylib"; fi

[windows]
clean:
    if (Test-Path dist) { Remove-Item -Recurse -Force dist }
    if (Test-Path build) { Remove-Item -Recurse -Force build }

[unix]
clean:
    rm -rf dist build
