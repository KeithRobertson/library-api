package services

import "testing"

func TestBookService_GetAllBooks(t *testing.T) {
	service := NewBookService()
	books, err := service.GetAllBooks()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(books) != 3 {
		t.Errorf("expected 3 books, got %d", len(books))
	}

	for _, book := range books {
		if book.Title == "" || book.Author == "" {
			t.Error("books should have title and author")
		}
	}
}
