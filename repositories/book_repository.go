package repositories

import "library-api/models"

type BookRepository interface {
	GetAll() ([]models.Book, error)
}
