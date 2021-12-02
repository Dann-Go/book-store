package domain

import "github.com/lib/pq"

//Book ...
type Book struct {
	ID      int      `db:"id" json:"id"`
	Title   string   `db:"title" json:"title"`
	Authors pq.StringArray `db:"authors"  json:"authors"`
	Year    string   `db:"year" json:"year"`
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
