package handler

import (
	"net/http"

	categoryHttp "bitgroup-cs/handler/v1/category/http"
	productHttp "bitgroup-cs/handler/v1/product/http"
	"bitgroup-cs/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ServiceUseCases struct {
	CategoryUseCase usecase.ICategoryUseCase
	ProductUseCase  usecase.IProductUseCase
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func GetRouters(uc ServiceUseCases) *gin.Engine {
	r := gin.Default()

	r.Use(corsMiddleware())

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "healthy",
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		// Categories endpoints
		categoryHandler := categoryHttp.NewHandler(uc.CategoryUseCase)
		categoriesGroup := apiV1.Group("/categories")
		{
			categoriesGroup.GET("", categoryHandler.FetchAll)
			categoriesGroup.GET("/:id", categoryHandler.FindById)
			categoriesGroup.POST("", categoryHandler.Create)
			categoriesGroup.PUT("/:id", categoryHandler.Update)
			categoriesGroup.DELETE("/:id", categoryHandler.Delete)
		}

		// Products endpoints
		productHandler := productHttp.NewHandler(uc.ProductUseCase)
		productsGroup := apiV1.Group("/products")
		{
			productsGroup.GET("", productHandler.FetchAll)
			productsGroup.GET("/:id", productHandler.FindById)
			productsGroup.POST("", productHandler.Create)
			productsGroup.PUT("/:id", productHandler.Update)
			productsGroup.DELETE("/:id", productHandler.Delete)
		}
	}

	return r
}
