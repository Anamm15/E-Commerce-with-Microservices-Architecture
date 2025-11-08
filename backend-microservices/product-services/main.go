package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"product-services/internal/configs"
	"product-services/internal/controllers"
	// "product-services/internal/controllers"
	"product-services/internal/repositories"
	"product-services/internal/services"
	"product-services/internal/storages"
	pb "product-services/pb"

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

	db := configs.ConnectDatabase()
	categoryRepository := repositories.NewCategoryRepository(db)
	productRepository := repositories.NewProductRepository(db)
	imageRepository := repositories.NewImageRepository(db)
	reviewRepository := repositories.NewReviewRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)
	productService := services.NewProductService(categoryRepository, productRepository, imageRepository, storage)
	reviewService := services.NewReviewService(reviewRepository)

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

	pb.RegisterCategoryServiceServer(grpcServer, categoryController)
	pb.RegisterProductServiceServer(grpcServer, productController)
	pb.RegisterReviewServiceServer(grpcServer, reviewController)

	// if err := migrations.Seeder(db); err != nil {
	// 	 log.Fatalf("error migration seeder: %v", err)
	// }

	log.Printf("🚀 gRPC server (ProductService) listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
