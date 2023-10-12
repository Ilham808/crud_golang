package model

import "github.com/stretchr/testify/mock"

// type UsersModelMock struct {
// 	mock.Mock
// }

// func (m *UsersModelMock) GetDatas() ([]Users, error) {
// 	args := m.Called()
// 	return args.Get(0).([]Users), args.Error(1)
// }

// type UsersModelMock struct {
// 	mock.Mock
// }

// func (um *UsersModelMock) GetDatas() ([]Users, error) {
// 	args := um.Called()
// 	return args.Get(0).([]Users), args.Error(1)
// }

type UsersModelMock struct {
	mock.Mock
}

func (m *UsersModelMock) GetDataById(id int) (*Users, error) {
	m.Mock.Called(id)
	return &Users{
		Name:     "Ilham Budiawan",
		Email:    "budiawanilham04@gmail.com",
		Password: "12345",
	}, nil
}

func (m *UsersModelMock) Delete(id int) error {
	m.Mock.Called(id)
	return nil
}

func (m *UsersModelMock) Update(users *Users) error {
	m.Mock.Called(users)
	return nil
}
