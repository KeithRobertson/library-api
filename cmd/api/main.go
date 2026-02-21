package main

import (
	"log"
	"net/http"

	"library-api/handlers"
	"library-api/middleware"
	"library-api/services"
)

func main() {
	bookService := services.NewBookService()
	bookHandler := handlers.NewBookHandler(bookService)

	mux := http.NewServeMux()
	mux.HandleFunc("/library/books", bookHandler.GetBooks)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.CORS(mux)))
}
