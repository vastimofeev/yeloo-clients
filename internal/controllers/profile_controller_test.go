package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"yeloo-clients/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProfileService для имитации ProfileService
type MockProfileService struct {
	mock.Mock
}

func (m *MockProfileService) GetProfile(id uint) (*models.Profile, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Profile), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProfileService) CreateProfile(profile *models.Profile) error {
	args := m.Called(profile)
	return args.Error(0)
}

func (m *MockProfileService) DeleteProfile(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetProfile(t *testing.T) {
	// Создаем моковый сервис
	mockService := new(MockProfileService)
	controller := &ProfileController{Service: mockService}

	// Настраиваем мок
	mockProfile := &models.Profile{
		ID:        1,
		FirstName: "John",
		// LastName:  "Doe",
		Birthdate: "1990-01-01",
		Email:     "john.doe@example.com",
	}
	mockService.On("GetProfile", uint(1)).Return(mockProfile, nil)

	// Создаем тестовый запрос
	router := gin.Default()
	router.GET("/profile/:id", controller.GetProfile)

	req, _ := http.NewRequest("GET", "/profile/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверяем результат
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John")
	mockService.AssertExpectations(t)
}

func TestCreateProfile(t *testing.T) {
	// Создаем моковый сервис
	mockService := new(MockProfileService)
	controller := &ProfileController{Service: mockService}

	// Настраиваем мок
	mockProfile := &models.Profile{
		ID:        1,
		FirstName: "Jane",
		// LastName:  "Doe",
		Birthdate: "1992-02-02",
		Email:     "jane.doe@example.com",
	}
	mockService.On("CreateProfile", mockProfile).Return(nil)

	// Создаем тестовый запрос
	router := gin.Default()
	router.POST("/profile", controller.CreateProfile)

	reqBody := `{"id":1,"first_name":"Jane","last_name":"Doe","birthdate":"1992-02-02","email":"jane.doe@example.com"}`
	req, _ := http.NewRequest("POST", "/profile", http.NoBody)
	req.Body = io.NopCloser(strings.NewReader(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверяем результат
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Jane")
	mockService.AssertExpectations(t)
}

func TestGetProfile_NotFound(t *testing.T) {
	// Создаем моковый сервис
	mockService := new(MockProfileService)
	controller := &ProfileController{Service: mockService}

	// Настраиваем мок
	mockService.On("GetProfile", uint(1)).Return(nil, fmt.Errorf("profile not found"))

	// Создаем тестовый запрос
	router := gin.Default()
	router.GET("/profile/:id", controller.GetProfile)

	req, _ := http.NewRequest("GET", "/profile/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверяем результат
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Profile not found")
	mockService.AssertExpectations(t)
}
