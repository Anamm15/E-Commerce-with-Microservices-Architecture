package grpc_clients

import (
	"log"

	orderpb "api-gateway/internal/pb/order"
	"google.golang.org/grpc"
)

type OrderClient struct {
	OrderClient orderpb.OrderServiceClient
}

func NewOrderClient(target string) (*OrderClient, error) {
	// Membuat koneksi ke server gRPC user service
	conn, err := grpc.Dial(target, grpc.WithInsecure()) // gunakan TLS di production
	if err != nil {
		return nil, err
	}

	// Membuat stub client dari hasil generate proto
	orderClient := orderpb.NewOrderServiceClient(conn)

	log.Printf("✅ Berhasil konek ke UserService di %s", target)

	return &OrderClient{
		OrderClient: orderClient,
	}, nil
}
