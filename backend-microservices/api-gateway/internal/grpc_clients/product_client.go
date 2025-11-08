package grpc_clients

import (
	"log"

	productpb "api-gateway/internal/pb/product"

	"google.golang.org/grpc"
)

// ProductServiceClientBundle menampung semua client dari product-services
type ProductServiceClient struct {
	ProductClient  productpb.ProductServiceClient
	CategoryClient productpb.CategoryServiceClient
	ReviewClient   productpb.ReviewServiceClient
	conn           *grpc.ClientConn
}

func NewProductServiceClient(target string) (*ProductServiceClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure()) // gunakan WithTransportCredentials() untuk TLS
	if err != nil {
		return nil, err
	}

	log.Printf("✅ Berhasil konek ke ProductService di %s", target)

	return &ProductServiceClient{
		ProductClient:  productpb.NewProductServiceClient(conn),
		CategoryClient: productpb.NewCategoryServiceClient(conn),
		ReviewClient:   productpb.NewReviewServiceClient(conn),
		conn:           conn,
	}, nil
}

// Close menutup koneksi gRPC ketika aplikasi berhenti
func (c *ProductServiceClient) Close() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
