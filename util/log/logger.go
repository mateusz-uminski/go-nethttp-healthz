package log

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(msg string, fields ...any)
	Infof(format string, args ...any)
	Error(err error, msg string, fields ...any)
	Errorf(err error, format string, args ...any)
	WithContext(ctx context.Context, keys ...any) Logger
}

type logger struct {
	sink zerolog.Logger
}

func Make(w io.Writer) logger {
	if w == nil {
		w = os.Stdout
	}
	return logger{
		sink: zerolog.New(w).With().Timestamp().Logger(),
	}
}

func (l logger) Info(msg string, fields ...any) {
	l.sink.Info().Fields(fields).Msg(msg)
}

func (l logger) Infof(format string, args ...any) {
	l.sink.Info().Msgf(format, args...)
}

func (l logger) Error(err error, msg string, fields ...any) {
	l.sink.Error().Err(err).Fields(fields).Msg(msg)
}

func (l logger) Errorf(err error, format string, args ...any) {
	l.sink.Error().Err(err).Msgf(format, args...)
}

func (l logger) WithContext(ctx context.Context, keys ...any) Logger {
	ctxSink := l.sink

	for _, key := range keys {
		if value, ok := ctx.Value(key).(string); ok {
			ctxSink = l.sink.With().Str(key.(string), value).Logger()
		}
	}

	return logger{sink: ctxSink}
}
