package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/wishes", getWishes)
	router.GET("/wishes/:id", getWishByID)

	router.Run("localhost:8080")
}
