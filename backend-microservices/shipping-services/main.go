package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"shipping-service/internal/configs"
	"shipping-service/internal/controllers"
	"shipping-service/internal/repositories"
	"shipping-service/internal/services"

	shippingpb "shipping-service/pb"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// checking environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️ Cannot load %s, using default .env", envFile)
		_ = godotenv.Load(".env")
	}

	// 🔹 Connecting to database
	db := configs.ConnectDatabase()

	// 🔹 Initializing repository, service, and controller
	shippingRepository := repositories.NewShippingRepository(db)
	shippingService := services.NewShippingService(shippingRepository)
	userController := controllers.NewShippingController(shippingService)

	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	shippingpb.RegisterShippingServiceServer(grpcServer, userController)

	// 🔹 Running gRPC listener
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "10004"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("gagal listen di port %s: %v", port, err)
	}

	log.Printf("✅ User Service berjalan di port %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gagal menjalankan gRPC server: %v", err)
	}
}
