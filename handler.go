package grpczap

import (
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var _ grpc.UnaryServerInterceptor = UnaryZapHandler

func UnaryZapHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	l := logger.With(zap.String("method", info.FullMethod))
	ctx = NewContext(ctx, l)
	return handler(ctx, req)
}
