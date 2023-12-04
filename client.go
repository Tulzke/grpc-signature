package grpc_signature

import "context"

func ClientName(ctx context.Context) string {
	name, ok := ctx.Value(ContextKey).(string)
	if !ok {
		return UnknownClient
	}

	return name
}
