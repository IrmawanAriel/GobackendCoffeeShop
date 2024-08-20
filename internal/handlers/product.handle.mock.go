package handlers

import (
	"testing"

	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	mockRepo := new(repositories.MockRepoProduct)
	product := &models.Product{Product_name: "Coffee", Price: 10000}

	mockRepo.On("CreateProduct", product).Return("1 data product created", nil)

	result, err := mockRepo.CreateProduct(product)

	assert.NoError(t, err)
	assert.Equal(t, "1 data product created", result)

	mockRepo.AssertExpectations(t)
}

func TestGetAllProduct(t *testing.T) {
	mockRepo := new(repositories.MockRepoProduct)
	products := &models.Products{
		{Product_name: "Coffee", Price: 10000},
		{Product_name: "Tea", Price: 15000},
	}

	pagination := &models.Pagination{Page: 1, Limit: 10}
	mockRepo.On("GetAllProduct", "", "", "", pagination).Return(products, nil)

	result, err := mockRepo.GetAllProduct("", "", "", pagination)

	assert.NoError(t, err)
	assert.Equal(t, products, result)

	mockRepo.AssertExpectations(t)
}

func TestGetProductById(t *testing.T) {
	mockRepo := new(repositories.MockRepoProduct)
	product := &models.Product{Product_name: "Coffee", Price: 10000}

	mockRepo.On("GetProductById", "1").Return(product, nil)

	result, err := mockRepo.GetProductById("1")

	assert.NoError(t, err)
	assert.Equal(t, product, result)

	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(repositories.MockRepoProduct)
	product := &models.Product{Product_name: "Coffee", Price: 12000}

	mockRepo.On("UpdateProduct", "1", product).Return("Product updated successfully", nil)

	result, err := mockRepo.UpdateProduct("1", product)

	assert.NoError(t, err)
	assert.Equal(t, "Product updated successfully", result)

	mockRepo.AssertExpectations(t)
}

func TestDeleteProductById(t *testing.T) {
	mockRepo := new(repositories.MockRepoProduct)

	mockRepo.On("DeleteProductById", "1").Return("Product deleted successfully", nil)

	result, err := mockRepo.DeleteProductById("1")

	assert.NoError(t, err)
	assert.Equal(t, "Product deleted successfully", result)

	mockRepo.AssertExpectations(t)
}

func TestGetFavoritesProduct(t *testing.T) {
	mockRepo := new(repositories.MockRepoProduct)
	products := &models.Products{
		{Product_name: "Coffee", Price: 10000},
		{Product_name: "Tea", Price: 15000},
	}

	mockRepo.On("GetFavoritesProduct", "1").Return(products, "", nil)

	result, msg, err := mockRepo.GetFavoritesProduct("1")

	assert.NoError(t, err)
	assert.Equal(t, "", msg)
	assert.Equal(t, products, result)

	mockRepo.AssertExpectations(t)
}

func TestAddFavoriteProduct(t *testing.T) {
	mockRepo := new(repositories.MockRepoProduct)

	mockRepo.On("AddFavoriteProduct", "1", "1").Return("Product added to favorite successfully", nil)

	result, err := mockRepo.AddFavoriteProduct("1", "1")

	assert.NoError(t, err)
	assert.Equal(t, "Product added to favorite successfully", result)

	mockRepo.AssertExpectations(t)
}

func TestDeleteFavoriteProduct(t *testing.T) {
	mockRepo := new(repositories.MockRepoProduct)

	mockRepo.On("DeleteFavoriteProduct", "1", "1").Return("Product deleted from favorite successfully", nil)

	result, err := mockRepo.DeleteFavoriteProduct("1", "1")

	assert.NoError(t, err)
	assert.Equal(t, "Product deleted from favorite successfully", result)

	mockRepo.AssertExpectations(t)
}
