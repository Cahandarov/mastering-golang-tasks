package handler

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"session-22/db"
	"session-22/model"
	"session-22/service"
	"session-22/util"
	"strconv"
)

type ApplicationHandler struct {
	BookService service.BookService
}

func NewApplicationHandler() *ApplicationHandler {
	log.Info("NewApplicationHandler generated")
	r := db.NewRepository(db.GetDB())
	s := service.NewBookService(r)
	app := &ApplicationHandler{
		s,
	}
	return app
}

func (a ApplicationHandler) getBooks(w http.ResponseWriter, r *http.Request) {
	log.Info("ActionLog.application.getBooks.start:", r.Method, r.URL)
	resp, err := a.BookService.GetBooks()
	util.PrepareResponse(w, resp, err)
	if err != nil {
		log.Error(err)
	}
	log.Info("ActionLog.application.getBooks.end")
}

func (a ApplicationHandler) getBookByID(w http.ResponseWriter, r *http.Request) {
	log.Info("ActionLog.application.getBookByID.start:", r.Method, r.URL)
	idString := r.PathValue("id")
	if idString == "" {
		http.Error(w, "Book ID not provided", http.StatusBadRequest)
		log.Error("Book ID not provided")
		return
	}
	bookID, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		log.Error("Invalid Book ID:", err)
		return
	}
	bookIDUint := uint(bookID)
	resp, serviceErr := a.BookService.GetBookByID(bookIDUint)
	if serviceErr != nil {
		if errors.Is(serviceErr, model.NotFoundError) {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		log.Error(serviceErr)
		return
	}

	util.PrepareResponse(w, resp, serviceErr)
	log.Info("ActionLog.application.getBookByID.end")
}

func (a ApplicationHandler) addBook(w http.ResponseWriter, r *http.Request) {
	log.Info("ActionLog.application.addBook.start:", r.Method, r.URL)

	var newBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := a.BookService.AddBook(newBook)
	var errorResponse *model.ErrorResponse
	if err != nil {
		errorResponse = &model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	util.PrepareResponse(w, nil, errorResponse)

	log.Info("ActionLog.application.addBook.end")
}

func (a ApplicationHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Info("ActionLog.application.deleteBook.start:", r.Method, r.URL)
	idString := r.PathValue("id")
	if idString == "" {
		http.Error(w, "Book ID not provided", http.StatusBadRequest)
		log.Error("Book ID not provided")
		return
	}
	bookID, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		log.Error("Invalid Book ID:", err)
		return
	}
	bookIDUint := uint(bookID)
	serviceErr := a.BookService.DeleteBook(bookIDUint)
	if serviceErr != nil {
		if errors.Is(serviceErr, model.NotFoundError) {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		log.Error(serviceErr)
		return
	}

	util.PrepareResponse(w, nil, serviceErr)
	log.Info("ActionLog.application.getBookByID.end")
}

func (a ApplicationHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	log.Info("ActionLog.application.updateBook.start:", r.Method, r.URL)
	idString := r.PathValue("id")
	if idString == "" {
		http.Error(w, "Book ID not provided", http.StatusBadRequest)
		log.Error("Book ID not provided")
		return
	}
	bookID, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		log.Error("Invalid Book ID:", err)
		return
	}
	bookIDUint := uint(bookID)

	var updatedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		util.PrepareResponse(w, model.InValidRequestError, &model.InValidRequestError)
		return
	}

	serviceErr := a.BookService.UpdateBook(bookIDUint, updatedBook)
	if serviceErr != nil {
		if errors.Is(serviceErr, model.NotFoundError) {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		log.Error(serviceErr)
		return
	}

	util.PrepareResponse(w, nil, serviceErr)
	log.Info("ActionLog.application.updateBook.end")
}
