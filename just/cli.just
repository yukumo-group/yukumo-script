build-cli: build-cli-win64 build-cli-win32 build-cli-linux64 build-cli-linux32 build-cli-macos
build-cli-debug: build-cli-win64-debug build-cli-win32-debug build-cli-linux64-debug build-cli-linux32-debug build-cli-macos-debug

# ---- release ----

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
