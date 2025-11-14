package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"product-services/internal/configs"
	"product-services/internal/controllers"
	// "product-services/internal/kafka"
	"product-services/internal/repositories"
	"product-services/internal/services"
	"product-services/internal/storages"

	productpb "product-services/pb"
	userpb "product-services/pb/user"

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

	ctx := context.Background()
	storage, err := storages.NewFirebaseStorage(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Firebase: %v", err)
	}

	//🔹Initializing connection to user service
	userServiceAddr := os.Getenv("USER_SERVICE_ADDR")
	if userServiceAddr == "" {
		userServiceAddr = "localhost:10001"
	}
	userConn, err := grpc.Dial(userServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)

	db := configs.ConnectDatabase()

	// kafkaProducer := kafka.NewKafkaProducer()
	// defer kafkaProducer.Close()

	categoryRepository := repositories.NewCategoryRepository(db)
	productRepository := repositories.NewProductRepository(db)
	imageRepository := repositories.NewImageRepository(db)
	reviewRepository := repositories.NewReviewRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)
	productService := services.NewProductService(categoryRepository, productRepository, imageRepository, storage)
	reviewService := services.NewReviewService(reviewRepository, userClient)

	categoryController := controllers.NewCategoryServer(categoryService)
	productController := controllers.NewProductServer(productService)
	reviewController := controllers.NewReviewServer(reviewService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10002"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	productpb.RegisterCategoryServiceServer(grpcServer, categoryController)
	productpb.RegisterProductServiceServer(grpcServer, productController)
	productpb.RegisterReviewServiceServer(grpcServer, reviewController)

	log.Printf("🚀 gRPC server (ProductService) listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
