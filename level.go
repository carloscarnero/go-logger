// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"log/slog"
	"strings"
)

// Level returns the requested format as an instance of [slog.Level], if it
// is a valid level name. Currently, only DEBUG, INFO, WARN, and ERROR are
// considered valid.
func Level(name string) (slog.Level, error) {
	n := strings.ToUpper(name)
	switch n {
	case "DEBUG":
		return slog.LevelDebug, nil

	case "INFO":
		return slog.LevelInfo, nil

	case "WARN":
		return slog.LevelWarn, nil

	case "ERROR":
		return slog.LevelError, nil

	default:
		return slog.LevelInfo, fmt.Errorf("%w: invalid level: %s", ErrLogger, name)
	}
}
