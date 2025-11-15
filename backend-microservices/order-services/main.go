package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"order-services/internal/configs"
	"order-services/internal/controllers"

	// "order-services/migrations"
	"order-services/internal/repositories"
	"order-services/internal/services"
	orderpb "order-services/pb/order"
	productpb "order-services/pb/product"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := fmt.Sprintf(".env.%s", env)

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️  Cannot load file %s, Trying .env default...", envFile)
		_ = godotenv.Load(".env")
	}

	productServiceAddr := os.Getenv("PRODUCT_SERVICE_ADDR")
	if productServiceAddr == "" {
		productServiceAddr = "localhost:10002"
	}
	productConn, err := grpc.Dial(productServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer productConn.Close()
	productClient := productpb.NewProductServiceClient(productConn)

	db := configs.ConnectDatabase()
	orderRepository := repositories.NewOrderRepository(db)
	orderItemRepository := repositories.NewOrderItemRepository(db)
	orderService := services.NewOrderService(orderRepository, orderItemRepository, productClient)
	orderController := controllers.NewOrderController(orderService)

	// if err := migrations.Seeder(db); err != nil {
	// 	log.Fatalf("error migration seeder: %v", err)
	// }

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, orderController)

	port := os.Getenv("ORDER_SERVICE_PORT")
	if port == "" {
		port = "10003"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen %s: %v", port, err)
	}

	log.Printf("✅ Server started on %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to run gRPC server: %v", err)
	}
}
