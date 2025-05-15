package helpers

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	TokenKey   = "authorization"
	ValidToken = "mysecrettoken123"
)

// Unary Client Interceptor
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = injectToken(ctx)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// Stream Client Interceptor
func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = injectToken(ctx)
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func injectToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, TokenKey, ValidToken)
}
