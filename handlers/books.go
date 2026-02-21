package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"library-api/services"
)

type BookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		log.Printf("error getting books: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Printf("error encoding books: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
