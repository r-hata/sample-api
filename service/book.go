package service

import (
	"github.com/r-hata/sample-api/model"
)

type BookService struct{}

func (BookService) Add(book *model.Book) error {
	_, err := DbEngine.Insert(book)
	if err != nil {
		return err
	}

	return nil
}

func (BookService) List() []model.Book {
	books := make([]model.Book, 0)
	err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&books)
	if err != nil {
		panic(err)
	}

	return books
}

func (BookService) Update(book *model.Book) error {
	_, err := DbEngine.Id(book.ID).Update(book)
	if err != nil {
		return err
	}

	return nil
}

func (BookService) Delete(id int) error {
	book := new(model.Book)
	_, err := DbEngine.Id(id).Delete(book)
	if err != nil {
		return err
	}

	return nil
}
