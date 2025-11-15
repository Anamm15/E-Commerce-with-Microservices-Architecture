package main

import (
	"log"

	"api-gateway/internal/grpc_clients"
	"api-gateway/internal/routes"
)

func main() {
	// 1️⃣ Inisialisasi koneksi ke service-service lewat gRPC
	userServiceClient, err := grpc_clients.NewUserClient("localhost:10001")
	if err != nil {
		log.Fatalf("Failed to connect to UserService: %v", err)
	}

	productClient, err := grpc_clients.NewProductServiceClient("localhost:10002")
	if err != nil {
		log.Fatalf("Failed to connect to ProductService: %v", err)
	}

	orderClient, err := grpc_clients.NewOrderClient("localhost:10003")
	if err != nil {
		log.Fatalf("Failed to connect to OrderService: %v", err)
	}

	// 2️⃣ Buat router Gin dan inject client-nya
	r := routes.SetupRouter(
		userServiceClient.UserClient,
		userServiceClient.AddressClient,
		productClient.CategoryClient,
		productClient.ProductClient,
		productClient.ReviewClient,
		orderClient.OrderClient,
	)

	// 3️⃣ Jalankan API Gateway
	log.Println("API Gateway berjalan di port 10000 🚀")
	r.Run(":10000")
}
