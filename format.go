// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2026 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"strings"
)

// Format normalizes and validates a format string. Normalization will
// transform the parameter into uppercase, while validation will make sure
// that only JSON and TEXT are accepted.
//
// If validation succeeds, the normalized value will be returned and a nil
// value for the error. Otherwise, the result is the empty string and a
// non-nil error (variable since it will contain the requested format
// as-is.)
//
// Deprecated: this function will be removed from a future major version.
// [NormalizeFormat] should be used instead, which is functionally
// identical.
func Format(name string) (string, error) {
	return NormalizeFormat(name)
}

// NormalizeFormat normalizes and consequently validates a format string.
// Normalization will transform the parameter into uppercase, while
// validation will make sure that only JSON and TEXT are accepted.
//
// If validation succeeds, the normalized value will be returned and a nil
// value for the error. Otherwise, the result is the empty string and a
// non-nil error (variable since it will contain the requested format
// as-is.)
func NormalizeFormat(name string) (string, error) {
	f := strings.ToUpper(name)
	switch f {
	case "JSON", "TEXT":
		return f, nil

	default:
		return "", fmt.Errorf("invalid log format: %s", name)
	}
}
