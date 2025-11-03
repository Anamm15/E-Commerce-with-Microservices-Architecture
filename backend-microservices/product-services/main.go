package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"product-services/configs"
	"product-services/controllers"
	// "product-services/migrations"
	"product-services/repositories"
	"product-services/routes"
	"product-services/services"
	"product-services/storages"

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
	ctx := context.Background()
	storage, err := storages.NewFirebaseStorage(ctx)
	if err != nil {
		log.Fatalf("Gagal menginisialisasi Firebase Storage: %v", err)
	}

	var (
		db                 *gorm.DB                        = config.ConnectDatabase()
		categoryRepository repositories.CategoryRepository = repositories.NewCategoryRepository(db)
		productRepository  repositories.ProductRepository  = repositories.NewProductRepository(db)
		imageRepository    repositories.ImageRepository    = repositories.NewImageRepository(db)
		reviewRepository   repositories.ReviewRepository   = repositories.NewReviewRepository(db)

		categoryService services.CategoryService = services.NewCategoryService(categoryRepository)
		productService  services.ProductService  = services.NewProductService(categoryRepository, productRepository, imageRepository, storage)
		reviewService   services.ReviewService   = services.NewReviewService(reviewRepository)

		categoryController controllers.CategoryController = controllers.NewCategoryController(categoryService)
		productController  controllers.ProductController  = controllers.NewProductController(productService)
		reviewController   controllers.ReviewController   = controllers.NewReviewController(reviewService)
	)

	routes.ProductRoutes(server, productController)
	routes.CategoryRoutes(server, categoryController)
	routes.ReviewRoutes(server, reviewController)
	// if err := migrations.Seeder(db); err != nil {
	// 	log.Fatalf("error migration seeder: %v", err)
	// }

	server.Run(":10000")
}
