package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/motorheads/user_service/controller"
)

func New() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	api.Use()
	{
		api.GET("/users", controller.GetAllUsers)
		api.GET("/user", controller.GetUser)
		api.POST("/user", controller.CreateUser)
		api.DELETE("/user", controller.DeleteUser)
		api.PUT("/user", controller.UpdateUser)
	}
	return router
}
