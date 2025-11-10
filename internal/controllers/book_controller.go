package controllers

import (
	"net/http"
	"strconv"

	"go-gin-starter/internal/services"
	"github.com/gin-gonic/gin"
)

// BookController handles the HTTP requests for books
type BookController struct {
	bookService *services.BookService
}

// NewBookController creates a new BookController
func NewBookController() *BookController {
	return &BookController{
		bookService: services.NewBookService(),
	}
}

// CreateBookRequest represents the request body for creating a book
type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// CreateBook handles the request to create a book
func (c *BookController) CreateBook(ctx *gin.Context) {
	var req CreateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := c.bookService.CreateBook(req.Title, req.Author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create book"})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

// GetBooks handles the request to get all books
func (c *BookController) GetBooks(ctx *gin.Context) {
	books, err := c.bookService.GetBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get books"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

// GetBook handles the request to get a book by ID
func (c *BookController) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	book, err := c.bookService.GetBook(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// UpdateBookRequest represents the request body for updating a book
type UpdateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBook handles the request to update a book
func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	var req UpdateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := c.bookService.UpdateBook(uint(id), req.Title, req.Author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update book"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// DeleteBook handles the request to delete a book
func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	err = c.bookService.DeleteBook(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete book"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}
