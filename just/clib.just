build-clib: build-clib-win64 build-clib-win32 build-clib-linux64 build-clib-linux32 build-clib-macos
build-clib-debug: build-clib-win64-debug build-clib-win32-debug build-clib-linux64-debug build-clib-linux32-debug build-clib-macos-debug

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
