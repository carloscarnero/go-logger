// logger: simple and opinionated log/Slog.Logger instance creator
// Copyright 2024-2025 by authors and contributors (see AUTHORS file)

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

func TestNew_valid(t *testing.T) {
	// The following will loop over all the test outputs, valid format
	// names, and valid level names. It is, in other words, a
	// multiplication of test cases.
	for _, output := range outputs {
		for _, format := range formats_valid {
			for _, level := range levels_valid {
				t.Run(fmt.Sprintf("output=%s, format=%s level=%s", output.name, format, level.name), func(t *testing.T) {
					assert.NotPanics(t, func() {
						actual, err := logger.New(output.writer, format, level.name, true)

						require.NotNil(t, actual)

						require.NoError(t, err)
					})
				})
			}
		}
	}
}

func TestNew_invalid(t *testing.T) {
	// The following will loop over all the test outputs, invalid format
	// names, and valid level names. It is, in other words, a
	// multiplication of test cases. Note that the format name is given
	// first when calling the New function, which means that there will be
	// no need to test both invalid format names and invalid level names.
	for _, output := range outputs {
		for _, format := range formats_invalid {
			for _, level := range levels_valid {
				t.Run(fmt.Sprintf("output=%s, format=%s level=%s", output.name, format, level.name), func(t *testing.T) {
					assert.NotPanics(t, func() {
						actual, err := logger.New(output.writer, format, level.name, true)

						require.Nil(t, actual)

						require.NotNil(t, err)
					})
				})
			}
		}
	}

	// The following will loop over all the test outputs, valid format
	// names, and invalid level names. It is, in other words, a
	// multiplication of test cases.
	for _, output := range outputs {
		for _, format := range formats_valid {
			for _, level := range levels_invalid {
				t.Run(fmt.Sprintf("output=%s, format=%s level=%s", output.name, format, level), func(t *testing.T) {
					assert.NotPanics(t, func() {
						actual, err := logger.New(output.writer, format, level, true)

						require.Nil(t, actual)

						require.NotNil(t, err)
					})
				})
			}
		}
	}
}

// The following variables, output, hold the io.Writer instances to use on
// the test functions. They exist as package-level variables because they
// will be used in other tests as well. Note that the variable names are
// snake cased just for symmetry with the test function names.

var outputs []struct {
	name   string
	writer io.Writer
} = []struct {
	name   string
	writer io.Writer
}{
	{"os.Stdout", os.Stdout},
	{"os.Stderr", os.Stderr},
	{"io.Discard", io.Discard},
	{"nil", nil},
}

func ExampleNew_levelDebug() {
	// This logger is created so it will not include the current date and
	// time, which makes its output completely predictable.
	logger, _ := logger.New(os.Stdout, "TEXT", "DEBUG", false)

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
	// level=DEBUG msg="Rough winds do shake the darling buds of May,"
	// level=INFO msg="And summer's lease hath all too short a date;"
	// level=WARN msg="Sometime too hot the eye of heaven shines,"
	// level=WARN msg="And often is his gold complexion dimm'd;"
	// level=INFO msg="And every fair from fair sometime declines,"
	// level=ERROR msg="By chance or nature's changing course untrimm'd;"
	// level=ERROR msg="But thy eternal summer shall not fade,"
	// level=INFO msg="Nor lose possession of that fair thou ow'st;"
	// level=DEBUG msg="Nor shall death brag thou wander'st in his shade,"
	// level=INFO msg="When in eternal lines to time thou grow'st:"
	// level=ERROR msg="   So long as men can breathe or eyes can see,"
	// level=DEBUG msg="   So long lives this, and this gives life to thee."
}

func ExampleNew_levelInfo() {
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

func ExampleNew_levelWarn() {
	// This logger is created so it will not include the current date and
	// time, which makes its output completely predictable.
	logger, _ := logger.New(os.Stdout, "TEXT", "WARN", false)

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
	// level=WARN msg="Sometime too hot the eye of heaven shines,"
	// level=WARN msg="And often is his gold complexion dimm'd;"
	// level=ERROR msg="By chance or nature's changing course untrimm'd;"
	// level=ERROR msg="But thy eternal summer shall not fade,"
	// level=ERROR msg="   So long as men can breathe or eyes can see,"
}

func ExampleNew_levelError() {
	// This logger is created so it will not include the current date and
	// time, which makes its output completely predictable.
	logger, _ := logger.New(os.Stdout, "TEXT", "ERROR", false)

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
	// level=ERROR msg="By chance or nature's changing course untrimm'd;"
	// level=ERROR msg="But thy eternal summer shall not fade,"
	// level=ERROR msg="   So long as men can breathe or eyes can see,"
}

func ExampleNew_levelNone() {
	// This logger is created so it will not include the current date and
	// time, which makes its output completely predictable.
	logger, _ := logger.New(os.Stdout, "TEXT", "NONE", false)

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

func ExampleNew_discard() {
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
