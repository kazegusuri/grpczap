package grpczap

import (
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var _ grpc.UnaryServerInterceptor = UnaryZapHandler
var _ grpc.StreamServerInterceptor = StreamZapHandler

func UnaryZapHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logger.With(zap.String("method", info.FullMethod))
	ctx = NewContext(ctx, l)
	return handler(ctx, req)
}

func StreamZapHandler(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	l := logger.With(zap.String("method", info.FullMethod))
	ctx = NewContext(ctx, l)
	return handler(ctx, stream)
}
