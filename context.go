package iamlog

import "context"

type key int

const (
	logContextKey key = iota
)

// WithContext returns a copy of context in which the log value is set.
func (l *zapLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey, l)
}

// FromContext returns the value of the log key on the ctx.
func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		l := ctx.Value(logContextKey)
		if l != nil {
			return l.(Logger)
		}
	}
	return WithName("Unknown-Context")
}
