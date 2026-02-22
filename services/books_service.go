package services

import (
	"library-api/models"
	"library-api/repositories"
)

type BookService interface {
	GetAllBooks() ([]models.Book, error)
}

type bookService struct {
	bookRepository repositories.BookRepository
}

func NewBookService(bookRepository repositories.BookRepository) BookService {
	return &bookService{bookRepository: bookRepository}
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return s.bookRepository.GetAll()
}
