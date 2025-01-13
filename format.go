// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2025 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"strings"
)

// Format normalizes the parameter transforming it to upper case, and
// validates it (currently, only JSON and TEXT are accepted.) If valid,
// this normalized value is returned as well as a nil error value. If
// invalid, an empty string will be returned along with a non nil error
// value (variable since it will contain the requested format as-is).
func Format(format string) (string, error) {
	f := strings.ToUpper(format)
	switch f {
	case "JSON", "TEXT":
		return f, nil

	default:
		return "", fmt.Errorf("invalid log format: %s", format)
	}
}
