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
// accepted as valid. In this case, the error value will be nil.
//
// If an invalid level name is passed, the returned level value is INFO,
// which is the zero value for that type.Regardless, it should probably be
// ignored since the error value will not be nil.
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
