package repositories

import (
	"go-gin-starter/internal/models"
	"go-gin-starter/pkg/database"
)

// BookRepository handles the database operations for the Book model
type BookRepository struct{}

// CreateBook creates a new book in the database
func (r *BookRepository) CreateBook(book *models.Book) error {
	return database.DB.Create(book).Error
}

// GetBooks retrieves all books from the database
func (r *BookRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := database.DB.Find(&books).Error
	return books, err
}

// GetBook retrieves a book by ID from the database
func (r *BookRepository) GetBook(id uint) (*models.Book, error) {
	var book models.Book
	err := database.DB.First(&book, id).Error
	return &book, err
}

// UpdateBook updates a book in the database
func (r *BookRepository) UpdateBook(book *models.Book) error {
	return database.DB.Save(book).Error
}

// DeleteBook deletes a book from the database
func (r *BookRepository) DeleteBook(book *models.Book) error {
	return database.DB.Delete(book).Error
}
