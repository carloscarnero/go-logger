// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024 by authors and contributors (see AUTHORS file)

package logger_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.carloscarnero.stream/logger"
)

func TestFormatValid(t *testing.T) {
	tcs := []struct {
		format string
	}{
		{"JSON"},
		{"json"},
		{"TEXT"},
		{"text"},
	}
	for _, tc := range tcs {
		t.Run(tc.format, func(t *testing.T) {
			expected := strings.ToUpper(tc.format)
			actual, err := logger.Format(tc.format)

			require.Equal(t, expected, actual)

			require.NoError(t, err)
		})
	}
}

func TestFormatInvalid(t *testing.T) {
	tcs := []struct {
		format string
	}{
		{""},
		{"YAML"},
		{"yaml"},
		{"SYSLOG"},
		{"syslog"},
		{"PLAIN"},
		{"plain"},
		{"XML"},
		{"xml"},
	}
	for _, tc := range tcs {
		t.Run(tc.format, func(t *testing.T) {
			expected := strings.ToUpper(tc.format)
			actual, err := logger.Format(tc.format)

			// In many cases, Go idiomatic practices expects the caller to
			// ignore the return value when an error value is not nil. This
			// test, however, makes sure that the internal implementation
			// still complies with the library's interface in this case.
			//
			// Instead of returning the zero value of a string, a design
			// decision has been made to return the upper cased
			// transformation of the input string.
			require.Equal(t, expected, actual)

			if assert.Error(t, err) {
				require.Equal(t, err, fmt.Errorf("%w: invalid format: %s", logger.ErrLogger, tc.format))
			}
		})
	}
}
