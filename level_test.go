// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2025 by authors and contributors (see AUTHORS file)

package logger_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.carloscarnero.stream/go-logger"
)

// The following variables, levels_valid and levels_invalid, hold the test
// cases, valid and invalid respectively. They exist as package-level
// variables because they will be used in other tests as well. Note that
// the variable names are snake cased just for symmetry with the test
// function names.
//
// The Level function under test can take any value as input, so the input
// space is basically unbounded. These test values below represent just
// strings that stress the documented requirements and expectations.

var levels_valid []struct {
	name  string
	level slog.Level
} = []struct {
	name  string
	level slog.Level
}{
	{"DEBUG", slog.LevelDebug},
	{"debug", slog.LevelDebug},
	{"deBUg", slog.LevelDebug},
	{"debUG", slog.LevelDebug},
	{"INFO", slog.LevelInfo},
	{"info", slog.LevelInfo},
	{"iNFo", slog.LevelInfo},
	{"inFO", slog.LevelInfo},
	{"WARN", slog.LevelWarn},
	{"warn", slog.LevelWarn},
	{"wARn", slog.LevelWarn},
	{"warN", slog.LevelWarn},
	{"ERROR", slog.LevelError},
	{"error", slog.LevelError},
	{"ErrOR", slog.LevelError},
	{"errOR", slog.LevelError},
}

var levels_invalid []string = []string{
	"",
	"INFORMATIONAL",
	"informational",
	"INFOrmational",
	"INFOrmaTIONal",
	"informATIOnal",
	"inforMATIONAL",
	"WARNING",
	"warning",
	"WArniNG",
	"warNING",
	"NONE",
	"none",
	"NOne",
	"NonE",
	"nONe",
	"noNE",
	"CRITICAL",
	"critical",
	"CRITical",
	"CRiticAL",
	"criTICal",
	"critiCAL",
	"TRACE",
	"trace",
	"TRacE",
	"trACe",
	"SPECIAL",
	"special",
	"SPeciAL",
	"spECIal",
}

func TestLevel_valid(t *testing.T) {
	for _, tc := range levels_valid {
		t.Run(fmt.Sprintf("level=%q", tc.name), func(t *testing.T) {
			expected := tc.level
			actual, err := logger.Level(tc.name)

			require.Equal(t, expected, actual)

			require.NoError(t, err)
		})
	}
}

func TestLevel_invalid(t *testing.T) {
	for _, tc := range levels_invalid {
		t.Run(fmt.Sprintf("level=%q", tc), func(t *testing.T) {
			expected := slog.LevelInfo
			actual, err := logger.Level(tc)

			// Except in some cases, hopefully properly documented, it is
			// expected the caller to ignore the return value when the
			// returned error is not nil. This test, however, makes sure
			// that the internal implementation still complies with the
			// library's interface in this case.
			//
			// The expected level is INFO, which happily coincides with the
			// zero value; however, this is actually enforced by the
			// implementation.
			require.Equal(t, expected, actual)

			if assert.Error(t, err) {
				require.Equal(t, err, fmt.Errorf("invalid log level: %s", tc))
			}
		})
	}
}
