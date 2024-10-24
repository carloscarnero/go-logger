// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"strings"
)

// Format returns the requested format as an upper cased string, if it is
// a valid format name. Currently, only JSON and TEXT are considered valid.
func Format(format string) (string, error) {
	f := strings.ToUpper(format)
	switch f {
	case "JSON", "TEXT":
		return f, nil

	default:
		return f, fmt.Errorf("%w: invalid format: %s", ErrLogger, format)
	}
}
