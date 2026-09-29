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
    if (Test-Path coverage.out) { Remove-Item -Force coverage.out }
    if (Test-Path coverage.html) { Remove-Item -Force coverage.html }
    if (Test-Path aqtk-shim/build) { Remove-Item -Recurse -Force aqtk-shim/build }
    foreach ($ver in 'aq1','aq2','aq10') { $dest = Join-Path 'third-party' $ver; if (Test-Path $dest) { Remove-Item -Recurse -Force $dest } }

[unix]
clean:
    rm -rf dist build coverage.out coverage.html aqtk-shim/build third-party/aq1 third-party/aq2 third-party/aq10
