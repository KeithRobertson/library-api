package services

import (
	"errors"
	"library-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) GetAll() ([]models.Book, error) {
	args := m.Called()
	return args.Get(0).([]models.Book), args.Error(1)
}

func TestBookService_GetAllBooks_returnsBooksFromRepository(t *testing.T) {
	mockBooksRepository := new(MockBookRepository)
	expectedBooks := []models.Book{
		{ID: "1", Title: "Book A"},
		{ID: "2", Title: "Book B"},
	}

	mockBooksRepository.
		On("GetAll").
		Return(expectedBooks, nil)

	service := NewBookService(mockBooksRepository)
	books, _ := service.GetAllBooks()

	assert.Equal(t, expectedBooks, books)
	mockBooksRepository.AssertExpectations(t)

}

func TestBookService_GetAllBooks_returnsErrorFromRepository(t *testing.T) {
	mockBooksRepository := new(MockBookRepository)

	expectedError := errors.New("Repository failure")

	mockBooksRepository.
		On("GetAll").
		Return([]models.Book{}, expectedError)

	service := NewBookService(mockBooksRepository)
	_, e := service.GetAllBooks()

	assert.Error(t, e)
	assert.Equal(t, expectedError, e)
	mockBooksRepository.AssertExpectations(t)
}
