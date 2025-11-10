package services

import (
	"go-gin-starter/internal/models"
	"go-gin-starter/internal/repositories"
)

// BookService handles the business logic for books
type BookService struct {
	bookRepository *repositories.BookRepository
}

// NewBookService creates a new BookService
func NewBookService() *BookService {
	return &BookService{
		bookRepository: &repositories.BookRepository{},
	}
}

// CreateBook handles the logic for creating a book
func (s *BookService) CreateBook(title, author string) (*models.Book, error) {
	book := &models.Book{
		Title:  title,
		Author: author,
	}

	err := s.bookRepository.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// GetBooks handles the logic for getting all books
func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.bookRepository.GetBooks()
}

// GetBook handles the logic for getting a book by ID
func (s *BookService) GetBook(id uint) (*models.Book, error) {
	return s.bookRepository.GetBook(id)
}

// UpdateBook handles the logic for updating a book
func (s *BookService) UpdateBook(id uint, title, author string) (*models.Book, error) {
	book, err := s.bookRepository.GetBook(id)
	if err != nil {
		return nil, err
	}

	book.Title = title
	book.Author = author

	err = s.bookRepository.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// DeleteBook handles the logic for deleting a book
func (s *BookService) DeleteBook(id uint) error {
	book, err := s.bookRepository.GetBook(id)
	if err != nil {
		return err
	}

	return s.bookRepository.DeleteBook(book)
}
