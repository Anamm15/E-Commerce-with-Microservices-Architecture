package grpc_clients

import (
	"log"

	shippingpb "api-gateway/internal/pb/shipping"
	"google.golang.org/grpc"
)

type ShippingServiceClient struct {
	ShippingClient shippingpb.ShippingServiceClient
}

func NewShippingClient(target string) (*ShippingServiceClient, error) {
	// Membuat koneksi ke server gRPC shipping service
	conn, err := grpc.Dial(target, grpc.WithInsecure()) // gunakan TLS di production
	if err != nil {
		return nil, err
	}

	// Membuat client stub dari proto
	ShippingClient := shippingpb.NewShippingServiceClient(conn)

	log.Printf("✅ Berhasil konek ke ShippingService di %s", target)

	return &ShippingServiceClient{
		ShippingClient: ShippingClient,
	}, nil
}
