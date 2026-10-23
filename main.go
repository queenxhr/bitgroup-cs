package main

import (
	"log"

	"bitgroup-cs/config"
	handler "bitgroup-cs/handler/v1"
	"bitgroup-cs/infrastructure/database"
	categoryRepo "bitgroup-cs/infrastructure/repository/category"
	productRepo "bitgroup-cs/infrastructure/repository/product"
	"bitgroup-cs/internal/usecase"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Database Connection (PostgreSQL: db_bitgroup_cs)
	dbConn, err := database.InitPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Fatal: Database initialization failed: %v", err)
	}

	// 3. Initialize Repositories (following logistic-service structure: model -> MySQL/Postgres)
	categoryRepository := categoryRepo.NewMysqlRepository(dbConn)
	productRepository := productRepo.NewMysqlRepository(dbConn)

	// 4. Initialize UseCases
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	productUseCase := usecase.NewProductUseCase(productRepository, categoryRepository)

	// 5. Initialize Router & HTTP Handlers
	router := handler.GetRouters(handler.ServiceUseCases{
		CategoryUseCase: categoryUseCase,
		ProductUseCase:  productUseCase,
	})

	// 6. Run HTTP Server
	log.Printf("Server is starting on port :%s...", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
