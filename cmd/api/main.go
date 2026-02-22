package main

import (
	"library-api/migrations"
	"library-api/repositories"
	"log"
	"net/http"

	"library-api/handlers"
	"library-api/middleware"
	"library-api/services"
)

func main() {
	migrations.RunMigrations()

	session, err := migrations.NewAppSession()
	if err != nil {
		log.Fatalf("cannot connect to keyspace: %v", err)
	}
	defer session.Close()

	bookRepository := repositories.NewBookRepository(session)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	mux := http.NewServeMux()

	api := http.NewServeMux()
	api.HandleFunc("/library/books", bookHandler.GetBooks)

	mux.Handle("/api/", http.StripPrefix("/api", api))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.CORS(mux)))
}
