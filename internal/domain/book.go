package domain

//Book ...
type Book struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    string   `json:"year"`
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
