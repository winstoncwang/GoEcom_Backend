package handlers

import (
	"ecommerce_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, models.Items)
}
