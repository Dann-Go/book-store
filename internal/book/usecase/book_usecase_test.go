package usecase

import (
	"github.com/Dann-Go/book-store/internal/domain"
	"github.com/Dann-Go/book-store/internal/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBookUsecase_Add(t *testing.T) {

	mockBookRepo := new(mocks.BookRepository)
	mockBook := domain.Book{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"Hello"},
		Year:    "2000-10-10",
	}
	t.Run("success", func(t *testing.T) {
		tmpMockBook := mockBook
		mockBookRepo.On("Add", mock.AnythingOfType("*domain.Book")).Return(nil).Once()

		ucase := NewBookUsecase(mockBookRepo)
		err := ucase.Add(&tmpMockBook)
		assert.NoError(t, err)
		assert.Equal(t, mockBook, tmpMockBook)
		mockBookRepo.AssertExpectations(t)
	})
}

func TestBookUsecase_GetAll(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBooks := make([]domain.Book, 0)
	mockBook := domain.Book{
		ID:      0,
		Title:   "",
		Authors: nil,
		Year:    "",
	}
	mockBooks = append(mockBooks, mockBook)

	t.Run("success", func(t *testing.T) {

		mockBookRepo.On("GetAll").Return(mockBooks, nil).Once()

		ucase := NewBookUsecase(mockBookRepo)
		a, err := ucase.GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockBookRepo.AssertExpectations(t)
	})
}

func TestBookUsecase_GetById(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := &domain.Book{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"Hello"},
		Year:    "2000-10-10",
	}
	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("GetById", mock.Anything).Return(mockBook, nil).Once()

		ucase := NewBookUsecase(mockBookRepo)
		a, err := ucase.GetById(mockBook.ID)
		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockBookRepo.AssertExpectations(t)
	})
}

func TestBookUsecase_Delete(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := domain.Book{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"Hello"},
		Year:    "2000-10-10",
	}
	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

		ucase := NewBookUsecase(mockBookRepo)
		err := ucase.Delete(mockBook.ID)
		assert.NoError(t, err)
		mockBookRepo.AssertExpectations(t)
	})
}

func TestBookUsecase_Update(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := domain.Book{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"Hello"},
		Year:    "2000-10-10",
	}
	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("Update", &mockBook, mock.AnythingOfTypeArgument("int")).Return(nil).Once()

		ucase := NewBookUsecase(mockBookRepo)
		err := ucase.Update(&mockBook, mockBook.ID)
		assert.NoError(t, err)
		mockBookRepo.AssertExpectations(t)
	})
}
