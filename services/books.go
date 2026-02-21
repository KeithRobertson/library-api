package services

import "library-api/models"

type BookService interface {
	GetAllBooks() ([]models.Book, error)
}

type bookService struct{}

func NewBookService() BookService {
	return &bookService{}
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return []models.Book{
		{Title: "The Go Programming Language", Author: "Alan Donovan"},
		{Title: "Clean Code", Author: "Robert Martin"},
		{Title: "The Pragmatic Programmer", Author: "Andy Hunt"},
	}, nil
}
