package handler

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"session-22/middleware"
)

func GetRouter() *http.ServeMux {
	log.Info("Getting ready routers on: /api/v1")

	app := NewApplicationHandler()
	auth := NewAuthHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/login", auth.login)
	mux.HandleFunc("GET /api/v1/books", app.getBooks)
	mux.HandleFunc("GET /api/v1/books/{id}", app.getBookByID)
	mux.HandleFunc("POST /api/v1/books", app.addBook)
	mux.Handle("DELETE /api/v1/books/{id}", middleware.CheckAuth(app.deleteBook))
	mux.Handle("PUT /api/v1/books/{id}", middleware.CheckAuth(app.updateBook))

	return mux
}
