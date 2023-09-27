package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blogs    []Blog `json:"-" gorm:"foreignkey:UserID"`
}

type UsersModel struct {
	db *gorm.DB
}

func (um *UsersModel) Init(db *gorm.DB) {
	um.db = db
}

func (s *UsersModel) GetDatas() ([]Users, error) {
	var users []Users
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UsersModel) GetDataById(id int) (*Users, error) {
	var user Users
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UsersModel) GetUserByEmail(email string) (*Users, error) {
	var user Users
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UsersModel) Create(user Users) error {
	err := s.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersModel) Update(id int, user *Users) (*Users, error) {

	existingUser := Users{}
	if err := s.db.Where("id = ?", id).First(&existingUser).Error; err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := s.db.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil

}

func (s *UsersModel) Delete(id int) error {
	var user Users
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
