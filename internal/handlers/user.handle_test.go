package handlers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUser(t *testing.T) {
	router := gin.Default()
	mockRepo := new(repositories.UserRepositoryMock)

	hashedPassword, err := pkg.HashPassword("testpassword")
	handler := NewUser(mockRepo)
	Users := models.Users{
		{Id: "13",
			Fullname: "hai",
			Email:    "hai@example.com",
			Password: hashedPassword,
			Role:     "admin",
		},
	}

	mockRepo.On("GetAllUser", mock.Anything).Return(&Users, nil)
	router.GET("/user", handler.FetchAll)

	req, err := http.NewRequest("GET", "/user", nil)
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "users fetched successfully", actualResponse.Message, "Response message does not match")
	assert.NotNil(t, actualResponse.Data, "Data does not match")
}

func TestGetUserById(t *testing.T) {
	router := gin.Default()
	mockRepo := new(repositories.UserRepositoryMock)

	hashedPassword, err := pkg.HashPassword("testpassword")

	handler := NewUser(mockRepo)
	User := models.User{
		Id:       "13",
		Fullname: "hai",
		Email:    "hai@example.com",
		Password: hashedPassword,
		Role:     "admin",
	}

	mockRepo.On("GetUserById", "1").Return(User, nil)
	router.GET("/user/:id", handler.FetchById)

	req, err := http.NewRequest("GET", "/user/1", nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")
	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "user fetched successfully", actualResponse.Message, "Response message does not match")
	assert.NotNil(t, actualResponse.Data, "Data does not match")
}

func TestRegister(t *testing.T) {
	router := gin.Default()
	mockRepo := new(repositories.UserRepositoryMock)

	handler := NewUser(mockRepo)
	mockRepo.On("InsertUser", mock.Anything).Return("User inserted successfully", nil)
	router.POST("/user/register", handler.Register)

	requestBody, _ := json.Marshal(map[string]string{
		"fullname": "haihaihai",
		"email":    "test@example.com",
		"password": "testpassword",
		"role":     "user",
	})

	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")
	assert.Equal(t, http.StatusCreated, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "account registered successfully", actualResponse.Message, "Response message does not match")
	assert.Equal(t, "User inserted successfully", actualResponse.Data, "Data does not match")
}

func TestLogin(t *testing.T) {
	router := gin.Default()
	mockRepo := new(repositories.UserRepositoryMock)
	hashedPassword, err := pkg.HashPassword("hashedPassword")
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	handler := NewUser(mockRepo)
	User := &models.User{
		Email:    "hai@example.com",
		Password: hashedPassword,
	}

	mockRepo.On("GetByEmail", mock.Anything).Return(User, nil)
	router.POST("/user/login", handler.Login)

	requestBody, _ := json.Marshal(map[string]string{
		"email":    "hai@example.com",
		"password": "hashedPassword",
	})

	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")
	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "login successful", actualResponse.Message, "Response message does not match")
	assert.NotNil(t, actualResponse.Data, "Data does not match")
}

func TestUpdateUserById(t *testing.T) {
	router := gin.Default()
	mockRepo := new(repositories.UserRepositoryMock)

	handler := NewUser(mockRepo)
	mockRepo.On("UpdateUser", "1", mock.Anything).Return("User updated successfully", nil)
	router.PATCH("/user/update/:id", handler.UpdateUserById)

	requestBody, _ := json.Marshal(map[string]string{
		"fullname": "John cena",
		"email":    "john.doe@example.com",
		"role":     "user",
		"password": "wanjaayy",
	})

	req, err := http.NewRequest("PATCH", "/user/update/1", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")
	assert.Equal(t, http.StatusCreated, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "User updated successfully", actualResponse.Message, "Response message does not match")
	assert.Equal(t, "User updated successfully", actualResponse.Data, "Data does not match")
}

func TestDeleteUser(t *testing.T) {
	router := gin.Default()
	userRepositoryMock := new(repositories.UserRepositoryMock)

	handler := NewUser(userRepositoryMock)

	userRepositoryMock.On("DeleteUserById", 1).Return("User deleted successfully", nil)
	router.DELETE("/user/delete/:id", handler.DeleteUser)

	req, err := http.NewRequest("DELETE", "/user/delete/1", nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, 200, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "user deleted successfully", actualResponse.Message, "Response message does not match")
	assert.Equal(t, "User deleted successfully", actualResponse.Data, "Status does not match")
}
