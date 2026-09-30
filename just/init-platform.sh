#!/usr/bin/env bash
# Build aqtk-shim Release for one platform and install it under third-party/<platform>/.
set -euo pipefail

platform="${1:?platform}"
shim="${2:?shim dir}"
prefix="${3:?install prefix}"

case "$platform" in
    win64) configure=x64; build=x64-release; dir=x64 ;;
    win32) configure=x86; build=x86-release; dir=x86 ;;
    linux64) configure=linux; build=linux-release; dir=linux ;;
    linux32) configure=linux32; build=linux32-release; dir=linux32 ;;
    macos) configure=macos; build=macos-release; dir=macos ;;
    *)
        echo "unknown platform '$platform' (win64, win32, linux64, linux32, macos)" >&2
        exit 1
        ;;
esac

os=$(uname -s)
case "$platform" in
    win64|win32)
        echo "platform '$platform' is built on Windows with MSVC" >&2
        exit 1
        ;;
    linux64|linux32)
        if [ "$os" != "Linux" ]; then
            echo "platform '$platform' is built on Linux" >&2
            exit 1
        fi
        ;;
    macos)
        if [ "$os" != "Darwin" ]; then
            echo "platform macos is built on macOS" >&2
            exit 1
        fi
        ;;
esac

(
    cd "$shim"
    cmake --preset "$configure" -DAQTK_ARCH="$platform" -DAQTK_STAGE_PHONTS=OFF -DCMAKE_INSTALL_PREFIX="$prefix"
    cmake --build --preset "$build"
)

rm -rf "$prefix/$platform"

cmake --install "$shim/build/$dir" --config Release --prefix "$prefix"
