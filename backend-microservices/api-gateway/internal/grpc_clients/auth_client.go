package grpc_clients

// import (
// 	authpb "api-gateway/proto/auth"
// 	"google.golang.org/grpc"
// )

// func NewAuthClient(addr string) (authpb.AuthServiceClient, error) {
// 	conn, err := grpc.Dial(addr, grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return authpb.NewAuthServiceClient(conn), nil
// }
