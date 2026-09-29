# Build CLI + clib for win64/win32, linux64/linux32, and macos (Zig cross).

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

test_tags := "test noaudio"
cover_pkgs := "./pkg/language/...,./pkg/utils/...,./pkg/characters,./pkg/phontsmanager,./pkg/api,./pkg/generator/tasks/singlesentence,./pkg/example"
