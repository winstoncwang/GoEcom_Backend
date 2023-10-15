package routes

import (
	"ecommerce_backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/items", handlers.GetItems)
	}
}
