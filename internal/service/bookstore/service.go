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
	AddBook(Book) (sql.Result, error)
	DeleteBook(int) (sql.Result, error)
}
type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (BookService, error) {
	return service{db, c}, nil
}

func (s service) AddBook(book Book) (sql.Result, error) {
	fmt.Printf("\nBody: %v\n\n", book)

	queryInsertBook := `INSERT INTO books (name, language, status, genre, editorial, author, publicado, price)
										VALUES (:name, :language, :status, :genre, :editorial, :author, :publicado, :price)`

	result, err := s.db.NamedExec(queryInsertBook, &book)

	// result, err := s.db.Exec(queryInsertBook, &book.Name, &book.Language, &book.Status, &book.Genre, &book.Editorial, &book.Author, &book.Publicado, &book.Price)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) FindByID(ID int) (*Book, error) {
	var book Book

	err := s.db.QueryRowx("SELECT * FROM books WHERE id=?", ID).StructScan(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s service) FindAll() []*Book {
	var list []*Book
	if err := s.db.Select(&list, "SELECT * FROM books"); err != nil {
		panic(err)
	}
	return list
}

func (s service) DeleteBook(ID int) (sql.Result, error) {

	result, err := s.db.Exec("DELETE FROM books WHERE id=?;", ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
