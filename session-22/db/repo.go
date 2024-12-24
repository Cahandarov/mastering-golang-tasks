package db

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"session-22/model"
)

type Repository interface {
	FindBooks() ([]model.Book, error)
	FindBookById(bookID uint) (*model.Book, error)
	Add(newBook model.Book) *model.ErrorResponse
	DeleteBook(bookID uint) error
	UpdateBook(bookID uint, updatedBook model.Book) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindBooks() ([]model.Book, error) {
	log.Info("ActionLog.repository.FindBooks.start")
	var books []model.Book

	err := r.db.Model(model.Book{}).Find(&books).Error

	if err != nil {
		log.Error("ActionLog.repository.FindBooks.error:", err)
		return books, err
	}
	log.Info("ActionLog.repository.FindBooks.end")

	return books, nil
}

func (r *repository) FindBookById(bookID uint) (*model.Book, error) {
	log.Info("ActionLog.repository.FindBookByID.start")
	var book model.Book

	err := r.db.Model(&model.Book{}).Where("id = ?", bookID).First(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("No matching record found for ID:", bookID)
			return nil, model.NotFoundError
		}
		log.Error("ActionLog.repository.FindBookByID.error:", err)
		return nil, err
	}

	log.Info("ActionLog.repository.FindBookByID.end")
	return &book, nil
}

func (r *repository) Add(newBook model.Book) *model.ErrorResponse {
	log.Info("ActionLog.repository.AddBookBy.start")

	err := r.db.Model(model.Book{}).Create(&newBook).Error

	if err != nil {
		log.Error("ActionLog.repository.AddBook.error:", err)
		return &model.UnexpectedError
	}
	log.Info("ActionLog.repository.AddBook.end")
	return nil
}

func (r *repository) DeleteBook(bookID uint) error {
	log.Info("ActionLog.repository.DeleteBook.start")

	result := r.db.Model(&model.Book{}).Where("id = ?", bookID).Delete(&model.Book{})
	if result.Error != nil {
		log.Error("ActionLog.repository.DeleteBook.error:", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Info("No matching record found for ID:", bookID)
		return model.NotFoundError
	}

	log.Info("ActionLog.repository.DeleteBook.end")
	return nil
}

func (r *repository) UpdateBook(bookID uint, updatedBook model.Book) error {
	log.Info("ActionLog.repository.UpdateBook.start")

	result := r.db.Model(&model.Book{}).Where("id = ?", bookID).Updates(updatedBook)

	if result.RowsAffected == 0 {
		log.Info("No matching record found for ID:", bookID)
		return model.NotFoundError
	}
	//if result.Error != nil {
	//	log.Error("ActionLog.repository.UpdateBook.error:", result.Error)
	//	return result.Error
	//}
	log.Info("ActionLog.repository.UpdateBook.end")
	return nil
}
