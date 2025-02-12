// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2025 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"log/slog"
	"strings"
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelNone  = slog.Level(256)
)

// Level normalizes and validates a level string. Normalization will
// transform the parameter into uppercase, while validation will make sure
// that only DEBUG, INFO, WARN, ERROR, and NONE are accepted. If validation
// succeeds, a corresponding instance of [slog.Level] will be returned and
// a nil value for the error; otherwise, the result is [slog.LevelInfo] and
// a non-nil error (variable since it will contain the requested level
// as-is.)
func Level(name string) (slog.Level, error) {
	n := strings.ToUpper(name)
	switch n {
	case "DEBUG":
		return LevelDebug, nil

	case "INFO":
		return LevelInfo, nil

	case "WARN":
		return LevelWarn, nil

	case "ERROR":
		return LevelError, nil

	case "NONE":
		return LevelNone, nil

	default:
		return LevelInfo, fmt.Errorf("invalid log level: %s", name)
	}
}
