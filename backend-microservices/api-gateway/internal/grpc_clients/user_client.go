package grpc_clients

import (
	"log"

	userpb "api-gateway/internal/pb/user"
	"google.golang.org/grpc"
)

type UserClient struct {
	Client userpb.UserServiceClient
}

func NewUserClient(target string) (*UserClient, error) {
	// Membuat koneksi ke server gRPC user service
	conn, err := grpc.Dial(target, grpc.WithInsecure()) // gunakan TLS di production
	if err != nil {
		return nil, err
	}

	// Membuat stub client dari hasil generate proto
	client := userpb.NewUserServiceClient(conn)

	log.Printf("✅ Berhasil konek ke UserService di %s", target)

	return &UserClient{
		Client: client,
	}, nil
}
