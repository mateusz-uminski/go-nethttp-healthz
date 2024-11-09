package fake

import (
	"context"

	"github.com/mateusz-uminski/go-nethttp-healthz/util/log"
)

type logger struct{}

func MakeLogger() logger {
	return logger{}
}

func (l logger) Info(msg string, fields ...any) {}

func (l logger) Infof(format string, args ...any) {}

func (l logger) Error(err error, msg string, fields ...any) {}

func (l logger) Errorf(err error, format string, args ...any) {}

func (l logger) WithContext(ctx context.Context, keys ...any) log.Logger {
	return logger{}
}
