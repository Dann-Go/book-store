package postegres

import (
	"github.com/Dann-Go/book-store/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "log"
)

type postgresqlRepository struct {
	Conn *sqlx.DB
}

func NewPostgresqlRepository(Conn *sqlx.DB) domain.BookRepository {
	return &postgresqlRepository{Conn}
}
func (p postgresqlRepository) Add(book *domain.Book) error {
	query := `INSERT INTO books(id, title, authors, year) VALUES ($1, $2, $3, $4);`
	_, err := p.Conn.Exec(query, book.ID, book.Title, book.Authors, book.Year)
	if err != nil {
		return err
	}

	return nil
}

func (p postgresqlRepository) GetAll() ([]domain.Book, error) {
	books := []domain.Book{}
	err := p.Conn.Select(&books, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (p postgresqlRepository) GetById(id int) (*domain.Book, error) {
	book := domain.Book{}
	err := p.Conn.Get(&book, "SELECT * FROM books where id=$1", id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (p postgresqlRepository) Delete(id int) error {
	query := `DELETE from books where id = $1;`
	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p postgresqlRepository) Update(book *domain.Book, id int) error {
	query := `UPDATE books set id = $2, title = $3, authors =$4, year = $5 where id = $1;`
	_, err := p.Conn.Exec(query, id, book.ID, book.Title, book.Authors, book.Year)
	if err != nil {
		return err
	}
	return nil
}
