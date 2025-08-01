package interceptor

import (
	"context"
	"errors"

	descAccess "github.com/UraharaKiska/go-chat-server/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type validator interface {
	Validate() error
}

type serviceProvider interface {
	AccessClient(ctx context.Context) descAccess.AccessV1Client
}

func ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if val, ok := req.(validator); ok {
		if err := val.Validate(); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}

func NewCheckPermissionInterceptor(sp serviceProvider) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// if servProv, ok := sp.(serviceProvider); ok {
			client := sp.AccessClient(ctx)
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, errors.New("metadata is not provided")
			}
			outgoingCtx := metadata.NewOutgoingContext(ctx, md)
			_, err := client.Check(outgoingCtx, &descAccess.CheckRequest{
				EndpointAddress: info.FullMethod,
			})
			if err != nil {
				return nil, err
			}
		// }
		return handler(ctx, req)
	}
}

// ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	
// 	return handler(ctx, req)
// }