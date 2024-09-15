package main

import (
	"/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// User registration and login routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.Run(":8080")
}
