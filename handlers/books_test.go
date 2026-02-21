package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"library-api/models"
)

type mockBookService struct {
	books []models.Book
	err   error
}

func (m *mockBookService) GetAllBooks() ([]models.Book, error) {
	return m.books, m.err
}

func TestGetBooksHandler_Success(t *testing.T) {
	mockService := &mockBookService{
		books: []models.Book{
			{Title: "Test Book", Author: "Test Author"},
		},
	}
	handler := NewBookHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/library/books", nil)
	w := httptest.NewRecorder()

	handler.GetBooks(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var books []models.Book
	if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(books) != 1 {
		t.Errorf("expected 1 book, got %d", len(books))
	}

	if books[0].Title != "Test Book" || books[0].Author != "Test Author" {
		t.Error("book should have correct title and author")
	}
}

func TestGetBooksHandler_ServiceError(t *testing.T) {
	mockService := &mockBookService{
		err: errors.New("database error"),
	}
	handler := NewBookHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/library/books", nil)
	w := httptest.NewRecorder()

	handler.GetBooks(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
}
