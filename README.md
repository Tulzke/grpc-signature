# Library for signing grpc requests

Allows you to understand on the server side which service made the request. 
This is implemented through a pair of interceptors. The client writes the name of the application to the metadata, 
and the server reads this data and writes it to the context.

## Usage

### How to use interceptors
#### On client side

```go
package main

import (
	"context"
	"github.com/tulzke/grpc_signature"
	"google.golang.org/grpc"
)

func main() {
	grpcClientDialOpts := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(
			grpc_signature.UnaryClientInterceptor("someApp"),
		),
		grpc.WithChainStreamInterceptor(
			grpc_signature.StreamClientInterceptor("someApp"),
		),
	}

	conn, err := grpc.DialContext(context.Background(), "", grpcClientDialOpts...)
}
```


#### On server side

```go
package main

import (
    grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/tulzke/grpc_signature"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_signature.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_signature.StreamServerInterceptor(),
		),
	)
}
```

### How to get the client name on server side

```go
package controller

import (
	"context"
	"examplev1"
	"github.com/tulzke/grpc_signature"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

func (s *Server) Example(ctx context.Context, r *examplev1.Request) (*examplev1.Response, error) {
	// First way
	clientName := grpc_signature.ClientName(ctx)
	
	// Second Way
	clientName = ctx.Value(grpc_signature.ContextKey).(string)
}
```