// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024 by authors and contributors (see AUTHORS file)

package logger_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.carloscarnero.stream/go-logger"
)

func TestNewLogger_valid(t *testing.T) {
	tcs := []struct {
		output io.Writer
		format string
		level  string
	}{
		{os.Stderr, "TEXT", "DEBUG"},
		{os.Stderr, "JSON", "DEBUG"},
		{os.Stderr, "TEXT", "INFO"},
		{os.Stderr, "JSON", "INFO"},
		{os.Stderr, "TEXT", "WARN"},
		{os.Stderr, "JSON", "WARN"},
		{os.Stderr, "TEXT", "ERROR"},
		{os.Stderr, "JSON", "ERROR"},

		{io.Discard, "TEXT", "DEBUG"},
		{io.Discard, "JSON", "DEBUG"},
		{io.Discard, "TEXT", "INFO"},
		{io.Discard, "JSON", "INFO"},
		{io.Discard, "TEXT", "WARN"},
		{io.Discard, "JSON", "WARN"},
		{io.Discard, "TEXT", "ERROR"},
		{io.Discard, "JSON", "ERROR"},

		{nil, "TEXT", "DEBUG"},
		{nil, "JSON", "DEBUG"},
		{nil, "TEXT", "INFO"},
		{nil, "JSON", "INFO"},
		{nil, "TEXT", "WARN"},
		{nil, "JSON", "WARN"},
		{nil, "TEXT", "ERROR"},
		{nil, "JSON", "ERROR"},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("format:%s,level:%s", tc.format, tc.level), func(t *testing.T) {
			assert.NotPanics(t, func() {
				actual, err := logger.New(tc.output, tc.format, tc.level, true)

				require.NotNil(t, actual)

				require.NoError(t, err)
			})
		})
	}
}

func TestNewLogger_invalid(t *testing.T) {
	tcs := []struct {
		output io.Writer
		format string
		level  string
	}{
		{os.Stderr, "", ""},
		{os.Stderr, "", "trace"},
		{os.Stderr, "", "CRITICAL"},
		{os.Stderr, "YAML", "warning"},
		{os.Stderr, "YAML", "TRACE"},
		{os.Stderr, "YAML", "INFORMATIONAL"},
		{os.Stderr, "yaml", "SPECIAL"},
		{os.Stderr, "yaml", "none"},
		{os.Stderr, "yaml", "informational"},
		{os.Stderr, "syslog", "warning"},
		{os.Stderr, "syslog", "trace"},
		{os.Stderr, "syslog", "informational"},
		{os.Stderr, "PLAIN", "TRACE"},
		{os.Stderr, "PLAIN", "none"},
		{os.Stderr, "PLAIN", "CRITICAL"},
		{os.Stderr, "plain", "INFORMATIONAL"},
		{os.Stderr, "plain", "SPECIAL"},
		{os.Stderr, "plain", "none"},
		{os.Stderr, "XML", "special"},
		{os.Stderr, "XML", "warning"},
		{os.Stderr, "XML", "TRACE"},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("format:%s,level:%s", tc.format, tc.level), func(t *testing.T) {
			assert.NotPanics(t, func() {
				actual, err := logger.New(tc.output, tc.format, tc.level, true)

				require.Nil(t, actual)

				require.NotNil(t, err)
			})
		})
	}
}

func ExampleNew_info() {
	// This logger is created so it will not include the current date and
	// time, which makes its output completely predictable.
	logger, _ := logger.New(os.Stdout, "TEXT", "INFO", false)

	logger.Info("Shall I compare thee to a summer's day?")
	logger.Info("Thou art more lovely and more temperate:")
	logger.Debug("Rough winds do shake the darling buds of May,")
	logger.Info("And summer's lease hath all too short a date;")
	logger.Warn("Sometime too hot the eye of heaven shines,")
	logger.Warn("And often is his gold complexion dimm'd;")
	logger.Info("And every fair from fair sometime declines,")
	logger.Error("By chance or nature's changing course untrimm'd;")
	logger.Error("But thy eternal summer shall not fade,")
	logger.Info("Nor lose possession of that fair thou ow'st;")
	logger.Debug("Nor shall death brag thou wander'st in his shade,")
	logger.Info("When in eternal lines to time thou grow'st:")
	logger.Error("   So long as men can breathe or eyes can see,")
	logger.Debug("   So long lives this, and this gives life to thee.")

	// Output:
	// level=INFO msg="Shall I compare thee to a summer's day?"
	// level=INFO msg="Thou art more lovely and more temperate:"
	// level=INFO msg="And summer's lease hath all too short a date;"
	// level=WARN msg="Sometime too hot the eye of heaven shines,"
	// level=WARN msg="And often is his gold complexion dimm'd;"
	// level=INFO msg="And every fair from fair sometime declines,"
	// level=ERROR msg="By chance or nature's changing course untrimm'd;"
	// level=ERROR msg="But thy eternal summer shall not fade,"
	// level=INFO msg="Nor lose possession of that fair thou ow'st;"
	// level=INFO msg="When in eternal lines to time thou grow'st:"
	// level=ERROR msg="   So long as men can breathe or eyes can see,"
}

func ExampleNew_disabled() {
	// This logger is created so it will not include the current date and
	// time, which makes its output completely predictable.
	logger, _ := logger.New(io.Discard, "TEXT", "DEBUG", false)

	logger.Info("Shall I compare thee to a summer's day?")
	logger.Info("Thou art more lovely and more temperate:")
	logger.Debug("Rough winds do shake the darling buds of May,")
	logger.Info("And summer's lease hath all too short a date;")
	logger.Warn("Sometime too hot the eye of heaven shines,")
	logger.Warn("And often is his gold complexion dimm'd;")
	logger.Info("And every fair from fair sometime declines,")
	logger.Error("By chance or nature's changing course untrimm'd;")
	logger.Error("But thy eternal summer shall not fade,")
	logger.Info("Nor lose possession of that fair thou ow'st;")
	logger.Debug("Nor shall death brag thou wander'st in his shade,")
	logger.Info("When in eternal lines to time thou grow'st:")
	logger.Error("   So long as men can breathe or eyes can see,")
	logger.Debug("   So long lives this, and this gives life to thee.")

	// Output:
}
