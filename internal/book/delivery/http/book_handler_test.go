package http

import (
	"github.com/Dann-Go/book-store/internal/domain"
	"github.com/Dann-Go/book-store/internal/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

//Ask tomorrow
//func TestBookHandler_Add(t *testing.T) {
//	mockBook := domain.Book{
//		ID:      0,
//		Title:   "Hello",
//		Authors: []string{"Hello"},
//		Year:    "2000-10-10",
//	}
//	tmpMockBook := mockBook
//	j, err := json.Marshal(tmpMockBook)
//
//	assert.Error(t, err)
//	mockUseCase := new(mocks.BookUsecase)
//	mockUseCase.On("Add", mock.AnythingOfType("*domain.Book")).Return(nil).Once()
//	handler := BookHandler{
//		BUsecase: mockUseCase,
//		valid:    validator.Validate{},
//	}
//	g := gin.New()
//	g.POST("/books", handler.Add)
//
//	w := httptest.NewRecorder()
//	req := httptest.NewRequest("POST", "/books", strings.NewReader(string(j)))
//
//	g.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//
//}

func TestBookHandler_Delete(t *testing.T) {
	mockBook := &domain.Book{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"Hello"},
		Year:    "2000-10-10",
	}
	mockUseCase := new(mocks.BookUsecase)
	mockUseCase.On("Delete", mock.Anything).Return(nil).Once()
	handler := BookHandler{
		BUsecase: mockUseCase,
		valid:    validator.Validate{},
	}
	g := gin.New()
	g.DELETE("/books/:id", handler.Delete)

	num := strconv.Itoa(mockBook.ID)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/books/:?id="+num, strings.NewReader(""))

	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBookHandler_GetAll(t *testing.T) {
	mockBooks := make([]domain.Book, 0)
	mockBook := domain.Book{
		ID:      0,
		Title:   "",
		Authors: nil,
		Year:    "",
	}
	mockBooks = append(mockBooks, mockBook)
	mockUseCase := new(mocks.BookUsecase)
	mockUseCase.On("GetAll").Return(mockBooks, nil).Once()
	handler := BookHandler{
		BUsecase: mockUseCase,
		valid:    validator.Validate{},
	}
	g := gin.New()
	g.GET("/books", handler.GetAll)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/books", strings.NewReader(""))

	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBookHandler_GetById(t *testing.T) {

	mockBook := &domain.Book{
		ID:      0,
		Title:   "",
		Authors: nil,
		Year:    "",
	}

	mockUseCase := new(mocks.BookUsecase)
	mockUseCase.On("GetById", mock.Anything).Return(mockBook, nil).Once()
	handler := BookHandler{
		BUsecase: mockUseCase,
		valid:    validator.Validate{},
	}
	g := gin.New()
	g.GET("/books/:id", handler.GetById)

	num := strconv.Itoa(mockBook.ID)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/books/:?id="+num, strings.NewReader(""))

	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

//Ask tomorrow
//func TestBookHandler_Update(t *testing.T) {
//	mockBook := &domain.Book{
//		ID:      0,
//		Title:   "Hello",
//		Authors: []string{"Hello"},
//		Year:    "2000-10-10",
//	}
//	j, err := json.Marshal(&mockBook)
//	assert.Error(t, err)
//	mockUseCase := new(mocks.BookUsecase)
//	mockUseCase.On("Update", &mockBook, mock.AnythingOfTypeArgument("int")).Return(nil).Once()
//	handler := BookHandler{
//		BUsecase: mockUseCase,
//		valid:    validator.Validate{},
//	}
//	g := gin.New()
//	g.PUT("/books/:id", handler.Update)
//
//	num := strconv.Itoa(mockBook.ID)
//	w := httptest.NewRecorder()
//	req := httptest.NewRequest("PUT", "/books/:?id="+num, strings.NewReader(string(j)))
//
//	g.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//}
