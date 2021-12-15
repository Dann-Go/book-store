package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dann-Go/book-store/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresqlRepository_Add(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	book := &domain.Book{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"hello"},
		Year:    "2000-01-01",
	}

	mock.ExpectExec("INSERT INTO books").WithArgs(book.Title, book.Authors, book.Year).WillReturnResult(sqlmock.NewResult(1, 1))
	br := NewPostgresqlRepository(sqlxDB)
	err = br.Add(book)
	assert.NoError(t, err)
	assert.Equal(t, 0, book.ID)
}

func TestPostgresqlRepository_GetAll(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	books := []domain.Book{{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"hello"},
		Year:    "2000-01-01",
	},
		{
			ID:      1,
			Title:   "Hello",
			Authors: []string{"hello"},
			Year:    "2000-01-01",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "title", "authors", "year"}).
		AddRow(books[0].ID, books[0].Title, books[0].Authors, books[0].Year).
		AddRow(books[1].ID, books[1].Title, books[1].Authors, books[1].Year)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	br := NewPostgresqlRepository(sqlxDB)
	res, err := br.GetAll()
	assert.NoError(t, err)
	assert.Len(t, res, 2)

}

func TestPostgresqlRepository_GetById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	books := &domain.Book{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"hello"},
		Year:    "2000-01-01",
	}
	rows := sqlmock.NewRows([]string{"id", "title", "authors", "year"}).
		AddRow(books.ID, books.Title, books.Authors, books.Year)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	br := NewPostgresqlRepository(sqlxDB)
	res, err := br.GetById(books.ID)
	assert.NoError(t, err)
	assert.Equal(t, res, books)

}

func TestPostgresqlRepository_Delete(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectExec("DELETE").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	br := NewPostgresqlRepository(sqlxDB)
	err = br.Delete(1)
	assert.NoError(t, err)

}

func TestPostgresqlRepository_Update(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	book := &domain.Book{
		ID:      0,
		Title:   "Hello",
		Authors: []string{"hello"},
		Year:    "2000-01-01",
	}

	mock.ExpectExec("UPDATE").WithArgs(book.ID, book.ID, book.Title, book.Authors, book.Year).WillReturnResult(sqlmock.NewResult(1, 1))
	br := NewPostgresqlRepository(sqlxDB)
	err = br.Update(book, book.ID)
	assert.NoError(t, err)

}
