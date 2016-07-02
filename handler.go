package grpczap

import (
	"github.com/kazegusuri/grpczap/zapctx"
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func UnaryZapHandler(logger zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		l := logger.With(zap.String("method", info.FullMethod))
		ctx = zapctx.NewContext(ctx, l)
		return handler(ctx, req)
	}
}
