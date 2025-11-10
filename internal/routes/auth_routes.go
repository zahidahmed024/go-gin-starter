package routes

import (
	"go-gin-starter/internal/controllers"
	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the authentication routes
func AuthRoutes(router *gin.Engine) {
	authController := controllers.NewAuthController()

	auth := router.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}
}
