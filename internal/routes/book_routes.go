package routes

import (
	"go-gin-starter/internal/controllers"
	"go-gin-starter/internal/middleware"
	"github.com/gin-gonic/gin"
)

// BookRoutes sets up the book routes
func BookRoutes(router *gin.Engine) {
	bookController := controllers.NewBookController()

	book := router.Group("/books")
	book.Use(middleware.AuthMiddleware())
	{
		book.POST("/", bookController.CreateBook)
		book.GET("/", bookController.GetBooks)
		book.GET("/:id", bookController.GetBook)
		book.PUT("/:id", bookController.UpdateBook)
		book.DELETE("/:id", bookController.DeleteBook)
	}
}
