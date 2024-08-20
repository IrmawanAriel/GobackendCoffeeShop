package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"IrmawanAriel/goBackendCoffeeShop/pkg"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/mock"
)

func TestGetAllProduct(t *testing.T) {
	router := gin.Default()
	mockRepo := new(repositories.MockRepoProduct)
	MockCLoud := new(pkg.MockCloudinary)

	handler := NewProduct(mockRepo, MockCLoud)
	rating := 5.0
	tanggal := time.Now()
	productImage := "image.png"
	products := models.Products{
		{
			Id:           "99",
			Description:  "makanan enak",
			Category:     "food",
			Stock:        3,
			Price:        35000,
			Product_name: "sayur duren",
			Rating:       &rating,
			Updated_at:   &tanggal,
			Image:        &productImage,
		},
	}

	mockRepo.On("GetAllProduct", mock.Anything).Return(&products, nil)
	router.GET("/product", handler.FetchAll)

	req, err := http.NewRequest("GET", "/product?limit=10&page=1", nil)
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Successfully retrieved products", actualResponse.Message, "Response message does not match")
	assert.NotNil(t, actualResponse.Data, "Data does not match")
}

func TestGetProductById(t *testing.T) {
	router := gin.Default()
	mockRepo := new(repositories.MockRepoProduct)
	MockCLoud := new(pkg.MockCloudinary)

	handler := NewProduct(mockRepo, MockCLoud)
	rating := 5.0
	tanggal := time.Now()
	productImage := "image.png"
	product := &models.Product{
		Id:           "99",
		Description:  "makanan enak",
		Category:     "food",
		Stock:        3,
		Price:        35000,
		Product_name: "sayur duren",
		Rating:       &rating,
		Updated_at:   &tanggal,
		Image:        &productImage,
	}

	mockRepo.On("GetProductById", "1").Return(product, nil)
	router.GET("/product/:id", handler.FetchById)

	req, err := http.NewRequest("GET", "/product/1", nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")
	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Successfully retrieved product", actualResponse.Message, "Response message does not match")
	assert.NotNil(t, actualResponse.Data, "Data does not match")
}

func TestDeleteProductById(t *testing.T) {
	router := gin.Default()
	mockRepo := new(repositories.MockRepoProduct)
	MockCLoud := new(pkg.MockCloudinary)

	handler := NewProduct(mockRepo, MockCLoud)

	mockRepo.On("DeleteProductById", "1").Return("Product deleted successfully", nil)
	router.DELETE("/product/:id", handler.DeleteProduct)

	req, err := http.NewRequest("DELETE", "/product/1", nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Product successfully deleted", actualResponse.Message, "Response message does not match")
	assert.Equal(t, "Product deleted successfully", actualResponse.Data, "Data does not match")
}
