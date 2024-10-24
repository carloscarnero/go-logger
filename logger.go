// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024 by authors and contributors (see AUTHORS file)

package logger

import (
	"errors"
	"io"
	"log/slog"
)

var ErrLogger = errors.New("can't create logger instance")

// New creates an instance of [slog.Logger] that sends its output to the
// given writer, using the given format, and with the specified level.
// Output lines include the event timestamp if the given parameter is set.
func New(output io.Writer, format string, level string, timestamps bool) (*slog.Logger, error) {
	if output == nil {
		output = io.Discard
	}

	f, err := Format(format)
	if err != nil {
		return nil, err
	}
	l, err := Level(level)
	if err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{
		Level: l,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 && !timestamps {
				return slog.Attr{}
			}
			return a
		},
	}
	switch f {
	case "JSON":
		return slog.New(slog.NewJSONHandler(output, opts)), nil

	case "TEXT":
		return slog.New(slog.NewTextHandler(output, opts)), nil

	default:
		panic("impossible condition")
	}
}
