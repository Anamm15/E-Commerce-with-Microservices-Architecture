package main

import (
	"log"

	"api-gateway/internal/grpc_clients"
	"api-gateway/internal/routes"
)

func main() {
	// 1️⃣ Inisialisasi koneksi ke service-service lewat gRPC
	// authClient, err := grpc_clients.NewAuthClient("localhost:50051")
	// if err != nil {
	// 	log.Fatalf("gagal konek ke auth service: %v", err)
	// }
	userClient, err := grpc_clients.NewUserClient("localhost:50052")
	if err != nil {
		log.Fatalf("gagal konek ke user service: %v", err)
	}

	// 2️⃣ Buat router Gin dan inject client-nya
	r := routes.SetupRouter(userClient.Client)

	// 3️⃣ Jalankan API Gateway
	log.Println("API Gateway berjalan di port 10000 🚀")
	r.Run(":10000")
}
