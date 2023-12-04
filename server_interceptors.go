package grpc_signature

import (
	"context"
	"goa.design/goa/v3/grpc/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const ContextKey = "client_name"

const UnknownClient = "unknown"

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		return handler(wrapServerContext(ctx), req)
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ss = middleware.NewWrappedServerStream(wrapServerContext(ss.Context()), ss)
		return handler(srv, ss)
	}
}

func wrapServerContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get(MetadataKey)
		if len(client) == 1 {
			ctx = context.WithValue(ctx, ContextKey, client[0])
			return ctx
		}
	}

	ctx = context.WithValue(ctx, ContextKey, UnknownClient)

	return ctx
}
