// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2026 by authors and contributors (see AUTHORS file)

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

var _levels = map[string]slog.Level{
	"DEBUG": LevelDebug,
	"INFO":  LevelInfo,
	"WARN":  LevelWarn,
	"ERROR": LevelError,
	"NONE":  LevelNone,
}

// Level normalizes and validates a level string. Normalization will
// transform the parameter into uppercase, while validation will make sure
// that only DEBUG, INFO, WARN, ERROR, and NONE are accepted.
//
// If validation succeeds, a corresponding instance of [slog.Level] will be
// returned and a nil value for the error; otherwise, the result is
// [slog.LevelInfo] and a non-nil error (variable since it will contain the
// requested level as-is.)
func Level(name string) (slog.Level, error) {
	n, err := NormalizeLevel(name)
	if err != nil {
		return LevelInfo, fmt.Errorf("invalid log level: %s", name)
	}

	return _levels[n], nil
}

// NormalizeLevel normalizes and validates a level string. Normalization
// will transform the parameter into uppercase, while validation will make
// sure that only DEBUG, INFO, WARN, ERROR, and NONE are accepted.
//
// If validation succeeds, the normalized value will be returned and a nil
// value for the error. Otherwise, the result is the empty string and a
// non-nil error (variable since it will contain the requested format
// as-is.)
func NormalizeLevel(name string) (string, error) {
	n := strings.ToUpper(name)
	switch n {
	case "DEBUG", "INFO", "WARN", "ERROR", "NONE":
		return n, nil

	default:
		return "", fmt.Errorf("invalid log level: %s", name)
	}
}
