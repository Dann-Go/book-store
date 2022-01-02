package domain

import (
	"github.com/lib/pq"
)

//Book ...
type Book struct {
	ID      int            `db:"id" json:"id" binding:"omitempty"`
	Title   string         `db:"title" json:"title" binding:"required,gte=1"`
	Authors pq.StringArray `db:"authors"  json:"authors" binding:"required,gte=1"`
	Year    string         `db:"year" json:"year" binding:"required,datetime=2006-01-02"`
}

// BookUsecase represent the book's usecase
type BookUsecase interface {
	Add(book *Book) error
	GetAll() ([]Book, error)
	GetById(id int) (*Book, error)
	GetByTitle(title string) ([]Book, error)
	Delete(id int) error
	Update(book *Book, id int) error
}

type BookRepository interface {
	Add(book *Book) error
	GetAll() ([]Book, error)
	GetById(id int) (*Book, error)
	GetByTitle(title string) ([]Book, error)
	Delete(id int) error
	Update(book *Book, id int) error
}
