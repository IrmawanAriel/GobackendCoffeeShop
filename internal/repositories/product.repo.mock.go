package repositories

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockRepoProduct struct {
	mock.Mock
}

func (m *MockRepoProduct) CreateProduct(data *models.Product) (string, error) {
	args := m.Called(data)
	return args.String(0), args.Error(1)
}

func (m *MockRepoProduct) GetAllProduct(search string, sort string, category string, pagination *models.Pagination) (*models.Products, error) {
	args := m.Called()
	return args.Get(0).(*models.Products), args.Error(1)
}

func (m *MockRepoProduct) GetProductById(id string) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockRepoProduct) UpdateProduct(id string, data *models.Product) (string, error) {
	args := m.Called(id, data)
	return args.String(0), args.Error(1)
}

func (m *MockRepoProduct) DeleteProductById(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *MockRepoProduct) GetFavoritesProduct(userID string) (*models.Products, string, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.Products), args.String(1), args.Error(2)
}

func (m *MockRepoProduct) AddFavoriteProduct(userId string, productId string) (string, error) {
	args := m.Called(userId, productId)
	return args.String(0), args.Error(1)
}

func (m *MockRepoProduct) DeleteFavoriteProduct(userId string, productId string) (string, error) {
	args := m.Called(userId, productId)
	return args.String(0), args.Error(1)
}
