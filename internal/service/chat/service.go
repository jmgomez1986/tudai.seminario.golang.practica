package chat

import (
	"github.com/jmoiron/sqlx"
	"tudai.seminario.golang.practica/internal/config"
)

// Book ...
type Book struct {
	ID        int
	Name      string
	Language  string
	Status    string
	Genre     string
	Editorial string
	Author    string
	Publicado string
	Price     string
}

// BookService ...
type BookService interface {
	FindAll()           []*Book
	FindByID(int)       *Book
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
	return nil
}

func (s service) FindAll() []*Book {
	var list []*Book
	if err := s.db.Select(&list, "SELECT * FROM book"); err != nil {
		panic(err)
	}
	return list
}
