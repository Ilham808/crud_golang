package model

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}

type BooksModel struct {
	db *gorm.DB
}

func (um *BooksModel) Init(db *gorm.DB) {
	um.db = db
}

func (s *BooksModel) GetDatas() ([]Books, error) {
	var books []Books
	err := s.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *BooksModel) GetDataById(id int) (*Books, error) {
	var book Books
	err := s.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BooksModel) Create(book Books) error {
	err := s.db.Save(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *BooksModel) Update(id int, book *Books) (*Books, error) {

	existingBook := Books{}
	if err := s.db.Where("id = ?", id).First(&existingBook).Error; err != nil {
		return nil, err
	}

	existingBook.Judul = book.Judul
	existingBook.Penulis = book.Penulis
	existingBook.Penerbit = book.Penerbit

	if err := s.db.Save(&existingBook).Error; err != nil {
		return nil, err
	}

	return &existingBook, nil

}

func (s *BooksModel) Delete(id int) error {
	var book Books
	if err := s.db.Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
