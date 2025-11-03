package main

import (
	"fmt"
	"log"
	"os"

	"order-services/configs"
	"order-services/controllers"
	// "order-services/migrations"
	"order-services/repositories"
	"order-services/routes"
	"order-services/services"

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
		db              *gorm.DB                     = config.ConnectDatabase()
		orderRepository repositories.OrderRepository = repositories.NewOrderRepository(db)
		orderService    services.OrderService        = services.NewOrderService(orderRepository)
		orderController controllers.OrderController  = controllers.NewOrderController(orderService)
	)
	routes.UserRoute(server, orderController)

	// if err := migrations.Seeder(db); err != nil {
	// 	log.Fatalf("error migration seeder: %v", err)
	// }

	server.Run(":10000")
}
