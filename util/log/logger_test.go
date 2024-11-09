package log_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/mateusz-uminski/go-nethttp-healthz/util/log"
	"github.com/rs/zerolog"

	"github.com/stretchr/testify/assert"
)

func TestLoggerInfo(t *testing.T) {
	t.Run("should log a plain message at info level without fields", func(t *testing.T) {
		var buffer bytes.Buffer
		logger := makeTestLogger(&buffer)
		msg := "logmessage"

		// when
		logger.Info(msg)

		// then
		expected := expectedInfoLog(msg)
		actual := acutalLog(&buffer)
		assert.Equal(t, expected, actual)
	})

	t.Run("should log a plain message at info level with fields", func(t *testing.T) {
		var buffer bytes.Buffer
		logger := makeTestLogger(&buffer)
		msg := "logmessage"

		// when
		logger.Info(msg, "example_field", "value")

		// then
		expected := expectedInfoLogWithField(`"example_field":"value"`, msg)
		actual := acutalLog(&buffer)
		assert.Equal(t, expected, actual)
	})

	t.Run("should log a formatted message at info level", func(t *testing.T) {
		var buffer bytes.Buffer
		logger := makeTestLogger(&buffer)
		msg := "logmessage"

		// when
		logger.Infof("%s", msg)

		// then
		expected := expectedInfoLog(msg)
		actual := acutalLog(&buffer)
		assert.Equal(t, expected, actual)
	})
}

func TestLoggerError(t *testing.T) {
	t.Run("should log a plain message at error level without fields", func(t *testing.T) {
		var buffer bytes.Buffer
		logger := makeTestLogger(&buffer)
		err := errors.New("exampleerror")
		msg := "logmessage"

		// when
		logger.Error(err, msg)

		// then
		expected := expectedErrorLog(err, msg)
		actual := acutalLog(&buffer)
		assert.Equal(t, expected, actual)
	})

	t.Run("should log a plain message at error level with fields", func(t *testing.T) {
		var buffer bytes.Buffer
		logger := makeTestLogger(&buffer)
		err := errors.New("exampleerror")
		msg := "logmessage"

		// when
		logger.Error(err, msg, "example_field", "value")

		// then
		expected := expectedErrorLogWithField(err, `"example_field":"value"`, msg)
		actual := acutalLog(&buffer)
		assert.Equal(t, expected, actual)
	})

	t.Run("should log a formatted message at error level", func(t *testing.T) {
		var buffer bytes.Buffer
		logger := makeTestLogger(&buffer)
		err := errors.New("exampleerror")
		msg := "logmessage"

		// when
		logger.Errorf(err, "%s", msg)

		// then
		expected := expectedErrorLog(err, msg)
		actual := acutalLog(&buffer)
		assert.Equal(t, expected, actual)
	})
}

func TestLoggerWithContext(t *testing.T) {
	t.Run("should not add context values if not present", func(t *testing.T) {
		var buffer bytes.Buffer
		logger := makeTestLogger(&buffer)
		ctx := context.Background()
		msg := "logmessage"

		// when
		l := logger.WithContext(ctx, "context_value_does_not_exist")
		l.Info(msg)

		// then
		expected := expectedInfoLog(msg)
		actual := acutalLog(&buffer)
		assert.Equal(t, expected, actual)
	})

	// t.Run("should add context values", func(t *testing.T) {
	// })
}

func fakeNow() time.Time {
	return time.Unix(0, 0).UTC() // January 1, 1970 at 00:00:00 UTC
}

func makeTestLogger(b *bytes.Buffer) log.Logger {
	zerolog.TimestampFunc = fakeNow
	return log.Make(b)
}

func acutalLog(b *bytes.Buffer) string {
	return strings.TrimSpace(b.String())
}

func expectedInfoLog(message string) string {
	timestamp := fakeNow().Format("1970-01-01T00:00:00Z")
	return fmt.Sprintf(`{"level":"info","time":"%s","message":"%s"}`, timestamp, message)
}

func expectedInfoLogWithField(field string, message string) string {
	timestamp := fakeNow().Format("1970-01-01T00:00:00Z")
	return fmt.Sprintf(`{"level":"info",%s,"time":"%s","message":"%s"}`, field, timestamp, message)
}

func expectedErrorLog(err error, message string) string {
	timestamp := fakeNow().Format("1970-01-01T00:00:00Z")
	return fmt.Sprintf(`{"level":"error","error":"%s","time":"%s","message":"%s"}`, err.Error(), timestamp, message)
}

func expectedErrorLogWithField(err error, field string, message string) string {
	timestamp := fakeNow().Format("1970-01-01T00:00:00Z")
	return fmt.Sprintf(`{"level":"error","error":"%s",%s,"time":"%s","message":"%s"}`, err.Error(), field, timestamp, message)
}
