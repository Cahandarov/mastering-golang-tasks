package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"session-22/db"
	"session-22/model"
)

type BookService interface {
	GetBooks() ([]model.BookResponse, *model.ErrorResponse)
	GetBookByID(bookID uint) (*model.BookResponse, *model.ErrorResponse)
	AddBook(newBook model.Book) *model.ErrorResponse
	DeleteBook(bookID uint) *model.ErrorResponse
	UpdateBook(bookID uint, updatedBook model.Book) *model.ErrorResponse
}

type bookService struct {
	repo db.Repository
}

func NewBookService(repo db.Repository) BookService {
	return &bookService{repo: repo}
}

func (r bookService) GetBooks() ([]model.BookResponse, *model.ErrorResponse) {
	log.Info("ActionLog.service.GetBooks.start")
	var bookResponse []model.BookResponse

	books, err := r.repo.FindBooks()

	if err != nil {
		log.Error("GetBooks.error: %v", err)
		return bookResponse, &model.UnexpectedError
	}

	for _, book := range books {
		bookResponse = append(bookResponse, model.BookResponse{
			Title:         book.Title,
			Author:        book.Author,
			Price:         book.Price,
			PublishedYear: book.PublishedYear,
		})
	}
	log.Info("ActionLog.service.GetBooks.end")
	return bookResponse, nil

}
func (r bookService) GetBookByID(bookID uint) (*model.BookResponse, *model.ErrorResponse) {
	log.Info("ActionLog.service.GetBookByID.start")
	var bookResponse model.BookResponse

	book, err := r.repo.FindBookById(bookID)
	if err != nil {
		log.Error("GetBookByID.error: %v", err)
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, &model.NotFoundError
		default:
			return nil, &model.UnexpectedError
		}

	}
	bookResponse = model.BookResponse{
		Title:         book.Title,
		Author:        book.Author,
		Price:         book.Price,
		PublishedYear: book.PublishedYear,
	}
	log.Info("ActionLog.service.GetBookByID.end")
	return &bookResponse, nil
}

func (r bookService) AddBook(newBook model.Book) *model.ErrorResponse {
	log.Info("ActionLog.service.GetBooks.start")

	err := r.repo.Add(newBook)

	if err != nil {
		log.Error("AddBooks.error: %v", err)
		return &model.UnexpectedError
	}
	log.Info("ActionLog.service.AddBook.end")
	return nil
}

func (r bookService) DeleteBook(bookID uint) *model.ErrorResponse {
	log.Info("ActionLog.service.GetBookByID.start")

	err := r.repo.DeleteBook(bookID)
	if err != nil {
		log.Error("DeleteBook.error: %v", err)
		switch {
		case errors.Is(err, model.NotFoundError):
			return &model.NotFoundError
		default:
			return &model.UnexpectedError
		}

	}
	log.Info("ActionLog.service.DeleteBook.end")
	return nil
}

func (r bookService) UpdateBook(bookID uint, updatedBook model.Book) *model.ErrorResponse {
	log.Info("ActionLog.service.UpdateBook.start")

	err := r.repo.UpdateBook(bookID, updatedBook)
	if err != nil {
		//log.Error("UpdateBook.error: %v", err)
		switch {
		case errors.Is(err, model.NotFoundError):
			return &model.NotFoundError
		default:
			return &model.UnexpectedError
		}

	}
	log.Info("ActionLog.service.UpdateBook.end")
	return nil
}
