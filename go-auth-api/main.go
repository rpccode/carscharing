package main

import (
	"go-auth-api/src/config"
	"go-auth-api/src/controllers"
	middlewares "go-auth-api/src/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDB()

	r := gin.Default()

	// Rutas públicas
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Rutas protegidas con autenticación
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/protected", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(200, gin.H{"message": "Bienvenido " + username})
		})
	}

	r.Run(":3000")
}
