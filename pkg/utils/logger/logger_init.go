//go:build !test

package logger

import (
	"io"
	"log/slog"
	"os"
)

var handler *slog.TextHandler
var syslogger *slog.Logger

// initialize logger
func init() {
	opts := &slog.HandlerOptions{AddSource: true}
	_, err := os.Stat("app.log")
	switch {
	case err == nil:
		// Directly open the file
		writer, err := os.OpenFile(
			"app.log",
			os.O_APPEND|os.O_WRONLY,
			0644,
		)
		if err != nil {
			panic(err)
		}
		opts.Level = slog.LevelInfo
		handler = slog.NewTextHandler(
			io.MultiWriter(writer, os.Stdout),
			opts,
		)
	case os.IsNotExist(err):
		_, err := os.Create("app.log")
		if err != nil {
			panic(err)
		}
		opts.Level = slog.LevelDebug
		writer, err := os.OpenFile(
			"app.log",
			os.O_APPEND|os.O_WRONLY,
			0644,
		)
		if err != nil {
			panic(err)
		}
		handler = slog.NewTextHandler(
			io.MultiWriter(writer, os.Stdout),
			opts,
		)
	default:
		panic(err)
	}
	syslogger = slog.New(handler)
}
