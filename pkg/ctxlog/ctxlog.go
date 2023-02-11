package ctxlog

import (
	"context"
	"go.uber.org/zap"
)

type ctxKey struct{}

func ContextFields(ctx context.Context) []zap.Field {
	fs, _ := ctx.Value(ctxKey{}).([]zap.Field)
	return fs
}

func WithFields(ctx context.Context, fields ...zap.Field) context.Context {
	if len(fields) == 0 {
		return ctx
	}

	return context.WithValue(ctx, ctxKey{}, mergeFields(ContextFields(ctx), fields))
}

// Ctx returns logger with additional fields from provided context.
func Ctx(ctx context.Context, l *zap.Logger) *zap.Logger {
	return l.With(ContextFields(ctx)...)
}

func mergeFields(a, b []zap.Field) []zap.Field {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	// NOTE: just append() here is unsafe. If a caller passed slice of fields
	// followed by ... with capacity greater than length, then simultaneous
	// logging will lead to a data race condition.
	//
	// See https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	c := make([]zap.Field, len(a)+len(b))
	n := copy(c, a)
	copy(c[n:], b)
	return c
}
