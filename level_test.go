// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024 by authors and contributors (see AUTHORS file)

package logger_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.carloscarnero.stream/go-logger"
)

func TestLevelValid(t *testing.T) {
	tcs := []struct {
		name  string
		level slog.Level
	}{
		{"DEBUG", slog.LevelDebug},
		{"debug", slog.LevelDebug},
		{"INFO", slog.LevelInfo},
		{"info", slog.LevelInfo},
		{"WARN", slog.LevelWarn},
		{"warn", slog.LevelWarn},
		{"ERROR", slog.LevelError},
		{"error", slog.LevelError},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.level
			actual, err := logger.Level(tc.name)

			require.Equal(t, expected, actual)

			require.NoError(t, err)
		})
	}
}

func TestLevelInvalid(t *testing.T) {
	tcs := []struct {
		name string
	}{
		{""},
		{"INFORMATIONAL"},
		{"informational"},
		{"WARNING"},
		{"warning"},
		{"NONE"},
		{"none"},
		{"CRITICAL"},
		{"critical"},
		{"TRACE"},
		{"trace"},
		{"SPECIAL"},
		{"special"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			expected := slog.LevelInfo
			actual, err := logger.Level(tc.name)

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
				require.Equal(t, err, fmt.Errorf("%w: invalid level: %s", logger.ErrLogger, tc.name))
			}
		})
	}
}
