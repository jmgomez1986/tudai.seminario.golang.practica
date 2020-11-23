package bookstore

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"tudai.seminario.golang.practica/internal/config"
)

// Book ...
type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Language  string `json:"language"`
	Status    string `json:"status"`
	Genre     string `json:"genre"`
	Editorial string `json:"editorial"`
	Author    string `json:"author"`
	Publicado string `json:"publicado"`
	Price     string `json:"price"`
}

// BookService ...
type BookService interface {
	FindAll() []*Book
	FindByID(int) (*Book, error)
	AddBook(Book) sql.Result
}
type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (BookService, error) {
	return service{db, c}, nil
}

func (s service) AddBook(book Book) sql.Result {
	fmt.Printf("\nBody: %v\n\n", book)

	queryInsertBook := `INSERT INTO book (name, language, status, genre, editorial, author, publicado, price)
										VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result := s.db.MustExec(queryInsertBook, &book.Name, &book.Language, &book.Status, &book.Genre, &book.Editorial, &book.Author, &book.Publicado, &book.Price)

	return result
}

func (s service) FindByID(ID int) (*Book, error) {
	var book Book

	err := s.db.QueryRowx("SELECT * FROM book WHERE id=?", ID).StructScan(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s service) FindAll() []*Book {
	var list []*Book
	if err := s.db.Select(&list, "SELECT * FROM book"); err != nil {
		panic(err)
	}
	return list
}
