//go:build test

package logger

import "log/slog"

// syslogger is nil under -tags test so Info/Error fall through to testing.T when set.
var syslogger *slog.Logger
