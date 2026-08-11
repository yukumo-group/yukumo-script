# AquesTalk2 native libraries

Place the Windows x64 AquesTalk2 binaries in `win64/`:

- `AquesTalk2.dll` (required at runtime)
- `AquesTalk2.lib` (optional; the C wrapper loads the DLL dynamically)

These files are gitignored. At runtime the process loads `third_party/aquestalk2/win64/AquesTalk2.dll` relative to the working directory (normally the repo root).
