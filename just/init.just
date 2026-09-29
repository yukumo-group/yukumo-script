shim_dir := justfile_directory() / "aqtk-shim"
stage_root := justfile_directory() / "third-party"
init_ps1 := justfile_directory() / "just" / "init-platform.ps1"
init_sh := justfile_directory() / "just" / "init-platform.sh"

# PLATFORM: win64, win32 (Windows/MSVC); linux64, linux32 (Linux); macos (macOS arm64).
# Install an aqtk-shim Release into third-party/aq{1,2,10}/PLATFORM.
[windows]
init platform:
    powershell.exe -NoProfile -File '{{init_ps1}}' '{{platform}}' '{{shim_dir}}' '{{stage_root}}'; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

[unix]
init platform:
    bash '{{init_sh}}' '{{platform}}' '{{shim_dir}}' '{{stage_root}}'
