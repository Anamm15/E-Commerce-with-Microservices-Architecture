package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"user-services/internal/configs"
	"user-services/internal/controllers"
	"user-services/internal/migrations"
	"user-services/internal/repositories"
	"user-services/internal/services"
	pb "user-services/pb"

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

	// 🔹 Running migration and seeder
	if err := migrations.Seeder(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
	}

	// 🔹 Initializing repository, service, and controller
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	pb.RegisterUserServiceServer(grpcServer, userController)

	// 🔹 Running gRPC listener
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "10001"
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
