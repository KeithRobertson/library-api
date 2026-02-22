package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"library-api/services"
)

type BookHandler interface {
	GetBooks(writer http.ResponseWriter, request *http.Request)
}

type bookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) BookHandler {
	return &bookHandler{bookService: bookService}
}

func (h *bookHandler) GetBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		log.Printf("error getting books: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(books); err != nil {
		log.Printf("error encoding books: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
