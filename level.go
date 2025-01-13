// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2025 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"log/slog"
	"strings"
)

// Level normalizes the parameter transforming it to upper case, and
// validates it (currently, only DEBUG, INFO, WARN, and ERROR are accepted.)
// If valid, a corresponding instance of [slog.Level] is returned as well
// as a nil error value. If invalid, [slog.LevelInfo] is returned along
// with a non nil error value (variable since it will contain the requested
// level as-is).
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
		return slog.LevelInfo, fmt.Errorf("invalid log level: %s", name)
	}
}
