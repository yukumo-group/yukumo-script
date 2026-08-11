# AquesTalk2 native libraries

Place the Windows x64 AquesTalk2 binaries in `win64/`:

- `AquesTalk2.dll`
- `AquesTalk2.lib`

These files are gitignored. The CGO wrapper in `pkg/generator/generatorwin` links against this directory.
