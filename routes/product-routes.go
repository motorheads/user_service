package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/motorheads/catalog_service/controller"
)

func New() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	api.Use()
	{
		api.GET("/products", controller.GetAllProducts)
		api.GET("/product", controller.GetProductByID)
		api.POST("/product", controller.CreateProduct)
		api.DELETE("/product", controller.DeleteProductByID)
		api.PUT("/product", controller.UpdateProductByID)
	}
	return router
}
