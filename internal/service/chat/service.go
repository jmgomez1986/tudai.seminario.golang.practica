package chat

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"tudai.seminario.golang.practica/internal/config"
)

// Book ...
type Book struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Language  string `db:"language"`
	Status    string `db:"status"`
	Genre     string `db:"genre"`
	Editorial string `db:"editorial"`
	Author    string `db:"author"`
	Publicado string `db:"publicado"`
	Price     string `db:"price"`
}

// BookService ...
type BookService interface {
	FindAll() []*Book
	FindByID(int) *Book
	AddBook(Book) error
}
type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (BookService, error) {
	return service{db, c}, nil
}

func (s service) AddBook(m Book) error {
	return nil
}

func (s service) FindByID(ID int) *Book {
	var book Book

	fmt.Printf("El id elegido es: %v\n", ID)

	err := s.db.QueryRowx("SELECT * FROM book WHERE id=?", ID).StructScan(&book)
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}

	return &book
}

func (s service) FindAll() []*Book {
	var list []*Book
	if err := s.db.Select(&list, "SELECT * FROM book"); err != nil {
		panic(err)
	}
	return list
}
