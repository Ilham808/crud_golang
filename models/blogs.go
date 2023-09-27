package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  uint   `json:"user" form:"user"`
}

type BlogModel struct {
	db *gorm.DB
}

func (um *BlogModel) Init(db *gorm.DB) {
	um.db = db
}

func (s *BlogModel) GetDatas() ([]Blog, error) {
	var blog []Blog
	err := s.db.Find(&blog).Error
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (s *BlogModel) GetDataById(id int) (*Blog, error) {
	var blog Blog
	err := s.db.Where("id = ?", id).First(&blog).Error
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (s *BlogModel) Create(blog Blog) error {
	err := s.db.Save(&blog).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *BlogModel) Update(id int, blog *Blog) (*Blog, error) {

	existingBlog := Blog{}
	if err := s.db.Where("id = ?", id).First(&existingBlog).Error; err != nil {
		return nil, err
	}

	existingBlog.Title = blog.Title
	existingBlog.Content = blog.Content

	if err := s.db.Save(&existingBlog).Error; err != nil {
		return nil, err
	}

	return &existingBlog, nil

}

func (s *BlogModel) Delete(id int) error {
	var blog Blog
	if err := s.db.Where("id = ?", id).First(&blog).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&blog).Error; err != nil {
		return err
	}

	return nil
}
