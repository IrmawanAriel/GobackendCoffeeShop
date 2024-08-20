package repositories

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUserById(id string) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetAllUser() (*models.Users, error) {
	args := m.Called()
	return args.Get(0).(*models.Users), args.Error(1)
}

func (m *UserRepositoryMock) UpdateUser(id string, data *models.User) (string, error) {
	args := m.Called(id, data)
	return args.String(0), args.Error(1)
}

func (m *UserRepositoryMock) InsertUser(data *models.User) (string, error) {
	args := m.Called(data)
	return args.String(0), args.Error(1)
}

func (m *UserRepositoryMock) DeleteUserById(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *UserRepositoryMock) GetByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) CreateUser(data *models.UserCreate) (string, error) {
	args := m.Called(data)
	return args.String(0), args.Error(1)
}
