package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"library-api/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBookService struct {
	mock.Mock
}

func (m *mockBookService) GetAllBooks() ([]models.Book, error) {
	args := m.Called()
	return args.Get(0).([]models.Book), args.Error(1)
}

func executeGetBooks(mockService *mockBookService) *httptest.ResponseRecorder {
	handler := NewBookHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()

	handler.GetBooks(w, req)
	return w
}

func TestGetBooksHandler_Success_successStatusCode(t *testing.T) {
	mockService := new(mockBookService)

	mockService.
		On("GetAllBooks").
		Return([]models.Book{}, nil)

	w := executeGetBooks(mockService)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)

}

func TestGetBooksHandler_Success_setsApplicationJsonContentType(t *testing.T) {
	mockService := new(mockBookService)

	mockService.
		On("GetAllBooks").
		Return([]models.Book{}, nil)

	w := executeGetBooks(mockService)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockService.AssertExpectations(t)

}

func TestGetBooksHandler_Success_noError(t *testing.T) {
	mockService := new(mockBookService)

	mockService.
		On("GetAllBooks").
		Return([]models.Book{}, nil)

	w := executeGetBooks(mockService)

	resp := w.Result()
	defer resp.Body.Close()

	var actual []models.Book
	err := json.NewDecoder(resp.Body).Decode(&actual)
	assert.NoError(t, err)

	mockService.AssertExpectations(t)

}

func TestGetBooksHandler_Success_writesBooks(t *testing.T) {
	mockService := new(mockBookService)
	expectedBooks := []models.Book{
		{ID: "1", Title: "Book A", Author: "Author A"},
		{ID: "2", Title: "Book B", Author: "Author B"},
	}
	mockService.
		On("GetAllBooks").
		Return(expectedBooks, nil)

	w := executeGetBooks(mockService)

	resp := w.Result()
	defer resp.Body.Close()
	var actual []models.Book
	json.NewDecoder(resp.Body).Decode(&actual)

	assert.Equal(t, expectedBooks, actual)
	mockService.AssertExpectations(t)

}
