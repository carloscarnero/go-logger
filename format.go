// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024 by authors and contributors (see AUTHORS file)

package logger

import (
	"fmt"
	"strings"
)

// Format returns the requested format name transformed to upper case, also
// as a string, if it is a valid format name. Currently, only be JSON and
// TEXT are accepted as valid. In this case, the error value will be nil.
//
// If an invalid format name is passed, the returned string is not the
// empty string, but its transformation into upper case, as if it were
// valid. Regardless, it should probably be ignored since the error value
// will not be nil.
func Format(format string) (string, error) {
	f := strings.ToUpper(format)
	switch f {
	case "JSON", "TEXT":
		return f, nil

	default:
		return f, fmt.Errorf("%w: invalid format: %s", ErrLogger, format)
	}
}
