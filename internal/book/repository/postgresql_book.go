package repository

import (
	"BookStore/internal/domain"
	"log"
)

var books = []domain.Book{{
	ID:      1,
	Title:   "Богатый папа, Бедный папа",
	Authors: []string{"Роберт Киосаки"},
	Year:    "1997",
},  {ID: 2,
	Title:   "Преступление и наказание",
	Authors: []string{"Фёдор Достоевский"},
	Year:    "1866"},
	{
	ID:      3,
	Title:   "Метро 2033",
	Authors: []string{"Дмитрий Глуховский"},
	Year:    "2002"}}

type postgresqlRepository struct {
	//db Connection
}

func NewPostgresqlRepository() domain.BookRepository {
	return &postgresqlRepository{}
}
func (p postgresqlRepository) Add(book *domain.Book) error {
	books = append(books, *book)
	return nil
}

func (p postgresqlRepository) GetAll() ([]domain.Book, error) {
	return books, nil
}

func (p postgresqlRepository) GetById(id int) (*domain.Book, error) {
	for i := range books {
		if books[i].ID == id {
			return &books[i], nil
		}
	}
	return nil, nil
}

func (p postgresqlRepository) Delete(id int) error {
	for i := range books {
		if books[i].ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	return nil
}

func (p postgresqlRepository) Update(book *domain.Book, id int) error {
	for i := range books {
		if books[i].ID == id {
			books[i] = *book
			break
		}
		log.Println("No such book was found. New book was added")
		books = append(books, *book)
	}
	return nil
}
