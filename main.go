package main

import (
	"21BDS0268_backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// User registration and login routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.Run(":8080")
}
