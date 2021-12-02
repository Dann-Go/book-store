package usecase

import (
	"github.com/Dann-Go/book-store/internal/domain"
	"time"
)

type bookUsecase struct {
	bookRepo   domain.BookRepository
	ctxTimeout time.Duration
}

func NewBookUsecase(bRepo domain.BookRepository, timeout time.Duration) domain.BookUsecase {
	return bookUsecase{
		bookRepo:   bRepo,
		ctxTimeout: timeout,
	}
}

func (b bookUsecase) Add(book *domain.Book) error {
	return b.bookRepo.Add(book)
}

func (b bookUsecase) GetAll() ([]domain.Book, error) {
	return b.bookRepo.GetAll()
}

func (b bookUsecase) GetById(id int) (*domain.Book, error) {
	return b.bookRepo.GetById(id)
}

func (b bookUsecase) Delete(id int) error {
	return b.bookRepo.Delete(id)
}

func (b bookUsecase) Update(book *domain.Book, id int) error {
	return b.bookRepo.Update(book, id)
}
