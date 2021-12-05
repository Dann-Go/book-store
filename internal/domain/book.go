package domain

import (
	"github.com/lib/pq"
	"time"
)

//Book ...
type Book struct {
	ID      int            `db:"id" json:"id" validate:"required"`
	Title   string         `db:"title" json:"title" validate:"required,gte=1"`
	Authors pq.StringArray `db:"authors"  json:"authors" validate:"required,gte=1"`
	Year    time.Time      `db:"year" json:"year" validate:"required"`
}

// BookUsecase represent the book's usecase
type BookUsecase interface {
	Add(book *Book) error
	GetAll() ([]Book, error)
	GetById(id int) (*Book, error)
	Delete(id int) error
	Update(book *Book, id int) error
}

type BookRepository interface {
	Add(book *Book) error
	GetAll() ([]Book, error)
	GetById(id int) (*Book, error)
	Delete(id int) error
	Update(book *Book, id int) error
}
