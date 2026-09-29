# Run non-GUI unit tests under tests/ (-tags test skips app.log; noaudio stubs playback).
test:
    go test -tags "{{test_tags}}" ./tests/...

# HTML coverage report for packages exercised by tests/ (coverage.out + coverage.html).
# Quote flags: PowerShell splits on commas and treats ".out" as a property access.
coverage:
    go test -tags "{{test_tags}}" ./tests/... "-coverpkg={{cover_pkgs}}" "-coverprofile=coverage.out"
    go tool cover "-html=coverage.out" "-o=coverage.html"
