package grpc_clients

import (
	"log"

	userpb "api-gateway/internal/pb/user"
	"google.golang.org/grpc"
)

type UserClient struct {
	UserClient    userpb.UserServiceClient
	AddressClient userpb.AddressServiceClient
}

func NewUserClient(target string) (*UserClient, error) {
	// Membuat koneksi ke server gRPC user service
	conn, err := grpc.Dial(target, grpc.WithInsecure()) // gunakan TLS di production
	if err != nil {
		return nil, err
	}

	// Membuat stub client dari hasil generate proto
	userClient := userpb.NewUserServiceClient(conn)
	addressClient := userpb.NewAddressServiceClient(conn)

	log.Printf("✅ Berhasil konek ke UserService di %s", target)

	return &UserClient{
		UserClient:    userClient,
		AddressClient: addressClient,
	}, nil
}
