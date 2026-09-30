# yukumo-script
#   just init <platform>   aqtk-shim Release -> third-party/<platform>/
#   just build             Go CLI + clib (just/clib.just, just/cli.just)

set windows-shell := ["powershell.exe", "-NoProfile", "-Command"]

import 'just/vars.just'
import 'just/test.just'
import 'just/clib.just'
import 'just/cli.just'
import 'just/engines.just'
import 'just/init.just'

default: build

build: build-clib build-cli
build-debug: build-clib-debug build-cli-debug
