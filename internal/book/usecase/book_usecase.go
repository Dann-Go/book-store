package usecase

import (
	"github.com/Dann-Go/book-store/internal/domain"
	"log"
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
	err := b.bookRepo.Add(book)
	if err != nil {
		return err
	}
	return nil
}

func (b bookUsecase) GetAll() ([]domain.Book, error) {
	result, err := b.bookRepo.GetAll()
	if err != nil {
		log.Println("No books where returned")
		return nil, err
	}
	return result, err
}

func (b bookUsecase) GetById(id int) (*domain.Book, error) {
	result, err := b.bookRepo.GetById(id)
	if err != nil {
		log.Println("No such book was found")
		return nil, err
	}
	return result, err
}

func (b bookUsecase) Delete(id int) error {
	err := b.bookRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (b bookUsecase) Update(book *domain.Book, id int) error {
	err := b.bookRepo.Update(book, id)
	if err != nil {
		return err
	}
	return nil
}
