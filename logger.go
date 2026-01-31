// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2026 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"io"
	"log/slog"
)

// New creates an instance of [slog.Logger] that sends its output to the
// given writer, with the given format, and at the specified level. Output
// lines include the event timestamp if the given parameter is set. If
// either an invalid format or level is requested, a nil is returned as
// well as an error (variable since it will contain the root cause.)
//
// Silent logging behavior can be achieved either by using a nil writer,
// [io.Discard], or by requesting the NONE level.
func New(output io.Writer, format string, level string, timestamps bool) (*slog.Logger, error) {
	f, err := Format(format)
	if err != nil {
		return nil, fmt.Errorf("can not create logger: %w", err)
	}
	l, err := Level(level)
	if err != nil {
		return nil, fmt.Errorf("can not create logger: %w", err)
	}

	// If no output is requested, then the discard handler is used, which
	// does not need to be further configured.
	if output == nil || output == io.Discard || l == LevelNone {
		return slog.New(slog.DiscardHandler), nil
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
	var handler slog.Handler
	switch f {
	case "JSON":
		handler = slog.NewJSONHandler(output, opts)

	case "TEXT":
		handler = slog.NewTextHandler(output, opts)
	}
	return slog.New(handler), nil
}
