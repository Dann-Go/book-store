package mocks

import (
	"github.com/Dann-Go/book-store/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

type BookUsecase struct {
	mock.Mock
}

func (m *BookUsecase) GetAll() ([]domain.Book, error) {
	ret := m.Called()

	var r0 []domain.Book
	if rf, ok := ret.Get(0).(func() []domain.Book); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]domain.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *BookUsecase) GetById(id int) (*domain.Book, error) {
	ret := m.Called()

	var r0 *domain.Book
	if rf, ok := ret.Get(0).(func(int) *domain.Book); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(*domain.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *BookUsecase) Delete(id int) error {
	ret := m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *BookUsecase) Update(book *domain.Book, id int) error {
	ret := m.Called(book, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Book, int) error); ok {
		r0 = rf(book, id)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

func (m *BookUsecase) Add(book *domain.Book) error {
	ret := m.Called(book)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Book) error); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
