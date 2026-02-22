package repositories

import (
	"library-api/models"

	"github.com/gocql/gocql"
)

type CassandraBookRepository struct {
	session *gocql.Session
}

func NewCassandraBookRepository(session *gocql.Session) *CassandraBookRepository {
	return &CassandraBookRepository{session: session}
}

func (r *CassandraBookRepository) GetAll() ([]models.Book, error) {
	iter := r.session.Query(`SELECT id, title, author FROM books`).Iter()

	var (
		id     gocql.UUID
		title  string
		author string
	)

	books := []models.Book{}

	for iter.Scan(&id, &title, &author) {
		books = append(books, models.Book{
			ID:     id.String(),
			Title:  title,
			Author: author,
		})
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return books, nil
}
