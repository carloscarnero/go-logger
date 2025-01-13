// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2025 by authors and contributors (see AUTHORS file)

package logger_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.carloscarnero.stream/go-logger"
)

// The following variables, formats_valid and formats_invalid, hold the
// test cases, valid and invalid respectively. They exist as package-level
// variables because they will be used in other tests as well. Note that
// the variable names are snake cased just for symmetry with the test
// function names.
//
// The Format function under test can take any value as input, so the input
// space is basically unbounded. These test values below represent just
// strings that stress the documented requirements and expectations.

var formats_valid []string = []string{
	"JSON",
	"json",
	"JSon",
	"jSOn",
	"jsoN",
	"TEXT",
	"text",
	"TExt",
	"tEXt",
	"texT",
}

var formats_invalid []string = []string{
	"",
	"YAML",
	"yaml",
	"YAml",
	"yAMl",
	"yamL",
	"SYSLOG",
	"syslog",
	"SYSlog",
	"sySLoG",
	"sysLOG",
	"PLAIN",
	"plain",
	"PLain",
	"Plain",
	"plAIn",
	"plaIN",
	"XML",
	"xml",
	"Xml",
	"xmL",
}

func TestFormat_valid(t *testing.T) {
	for _, tc := range formats_valid {
		t.Run(fmt.Sprintf("format=%q", tc), func(t *testing.T) {
			expected := strings.ToUpper(tc)
			actual, err := logger.Format(tc)

			require.Equal(t, expected, actual)

			require.NoError(t, err)
		})
	}
}

func TestFormat_invalid(t *testing.T) {
	for _, tc := range formats_invalid {
		t.Run(fmt.Sprintf("format=%q", tc), func(t *testing.T) {
			expected := ""
			actual, err := logger.Format(tc)

			require.Equal(t, expected, actual)

			if assert.Error(t, err) {
				require.Equal(t, err, fmt.Errorf("invalid log format: %s", tc))
			}
		})
	}
}
