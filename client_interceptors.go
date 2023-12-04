package grpc_signature

import (
	"context"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

const MetadataKey = "client-name"

func UnaryClientInterceptor(clientName string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = metadata.AppendToOutgoingContext(ctx, MetadataKey, clientName)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamClientInterceptor(clientName string) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, MetadataKey, clientName)
		return streamer(ctx, desc, cc, method, opts...)
	}
}
