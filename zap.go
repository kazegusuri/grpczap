package grpczap

import (
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
)

type zapKey struct{}

var logger zap.Logger

func init() {
	logger = zap.NewJSON()
}

func SetLogger(l zap.Logger) {
	logger = l
}

func NewContext(ctx context.Context, l zap.Logger) context.Context {
	return context.WithValue(ctx, zapKey{}, l)
}

func FromContext(ctx context.Context) (zap.Logger, bool) {
	l, ok := ctx.Value(zapKey{}).(zap.Logger)
	return l, ok
}

func MustFromContext(ctx context.Context) zap.Logger {
	l, ok := ctx.Value(zapKey{}).(zap.Logger)
	if !ok {
		panic("could not find zap.Logger from context")
	}
	return l
}
