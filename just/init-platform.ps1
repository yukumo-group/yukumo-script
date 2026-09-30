param(
    [Parameter(Mandatory = $true)][string]$Platform,
    [Parameter(Mandatory = $true)][string]$Shim,
    [Parameter(Mandatory = $true)][string]$Prefix
)

$ErrorActionPreference = 'Stop'

switch ($Platform) {
    'win64' { $configure = 'x64'; $build = 'x64-release'; $dir = 'x64' }
    'win32' { $configure = 'x86'; $build = 'x86-release'; $dir = 'x86' }
    'linux64' { $configure = 'linux'; $build = 'linux-release'; $dir = 'linux' }
    'linux32' { $configure = 'linux32'; $build = 'linux32-release'; $dir = 'linux32' }
    'macos' { $configure = 'macos'; $build = 'macos-release'; $dir = 'macos' }
    default {
        Write-Host "unknown platform '$Platform' (win64, win32, linux64, linux32, macos)"
        exit 1
    }
}

if ($Platform -ne 'win64' -and $Platform -ne 'win32') {
    Write-Host "platform '$Platform' is built on that host; Windows builds win64 and win32"
    exit 1
}

Push-Location $Shim
try {
    cmake --preset $configure -DAQTK_ARCH="$Platform" -DAQTK_STAGE_PHONTS=OFF "-DCMAKE_INSTALL_PREFIX=$Prefix"
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
    cmake --build --preset $build
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
}
finally {
    Pop-Location
}

$dest = Join-Path $Prefix $Platform
if (Test-Path $dest) {
    Remove-Item -Recurse -Force $dest
}

cmake --install (Join-Path $Shim "build/$dir") --config Release --prefix $Prefix
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
