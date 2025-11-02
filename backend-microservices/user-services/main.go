package main

import (
	"fmt"
	"log"
	"os"

	"user-services/configs"
	"user-services/controllers"
	"user-services/migrations"
	"user-services/repositories"
	"user-services/routes"
	"user-services/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
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

	server := gin.Default()

	var (
		db             *gorm.DB                    = config.ConnectDatabase()
		userRepository repositories.UserRepository = repositories.NewUserRepository(db)
		userService    services.UserService        = services.NewUserService(userRepository)
		userController controllers.UserController  = controllers.NewUserController(userService)
	)
	routes.UserRoute(server, userController)

	if err := migrations.Seeder(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
	}

	server.Run(":10000")
}
