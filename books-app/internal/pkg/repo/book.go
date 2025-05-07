package repo

import (
	"github.com/xavicci/gRPC-ModernAPI/books-app/internal/pkg/model"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func GetNewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) AddBook(book *model.DBBook) {
	br.db.Create(book)
}

func (br *BookRepository) UpdateBook(book *model.DBBook) {
	br.db.Model(&book).Where("isbn = ?", book.Isbn).Update("name", "publiser")
}

func (br *BookRepository) GetBook(isbn int) *model.DBBook {
	var book model.DBBook
	br.db.First(&book, isbn)
	return &book
}

func (br *BookRepository) GetAllBooks() ([]*model.DBBook, error) {
	books := make([]*model.DBBook, 0)
	err := br.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (br *BookRepository) RemoveBook(isbn int) {
	br.db.Delete(&model.DBBook{}, isbn)
}
