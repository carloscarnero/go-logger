// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2025 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"strings"
)

// Format normalizes and validates a format string. Normalization will
// transform the parameter into uppercase, while validation will make sure
// that only JSON and TEXT are accepted. If validation succeeds, the
// normalized value will be returned and a nil value for the error.
// Otherwise, the result is the empty string and a non-nil error (variable
// since it will contain the requested format as-is.)
func Format(format string) (string, error) {
	f := strings.ToUpper(format)
	switch f {
	case "JSON", "TEXT":
		return f, nil

	default:
		return "", fmt.Errorf("invalid log format: %s", format)
	}
}
