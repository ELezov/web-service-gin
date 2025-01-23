package main

import (
	"web-service-gin/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/wishes", service.GetWishes)
	router.GET("/wishes/:id", service.GetWishByID)

	router.Run("localhost:8080")
}
