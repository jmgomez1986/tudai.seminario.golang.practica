package bookstore

import (
	"database/sql"

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
	UpdateBook(int, Book) (sql.Result, error)
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

	sqlStatement := `INSERT INTO books (
													name,
													language,
													status,
													genre,
													editorial,
													author,
													publicado,
													price
												)
												VALUES (
													:name,
													:language,
													:status, :genre,
													:editorial,
													:author,
													:publicado,
													:price
												);`

	result, err := s.db.NamedExec(sqlStatement, &book)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) FindByID(ID int) (*Book, error) {
	var book Book
	sqlStatement := `SELECT * FROM books WHERE id=?;`

	err := s.db.QueryRowx(sqlStatement, ID).StructScan(&book)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s service) FindAll() []*Book {
	var list []*Book
	sqlStatement := `SELECT * FROM books;`

	if err := s.db.Select(&list, sqlStatement); err != nil {
		panic(err)
	}
	return list
}

func (s service) DeleteBook(ID int) (sql.Result, error) {

	sqlStatement := `DELETE FROM books WHERE id=?;`

	result, err := s.db.Exec(sqlStatement, ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s service) UpdateBook(ID int, book Book) (sql.Result, error) {

	book.ID = ID

	sqlStatement := `UPDATE books
										SET
											name=:name,
											language=:language,
											status=:status,
											genre=:genre,
											editorial=:editorial,
											author=:author,
											publicado=:publicado,
											price=:price
										WHERE id=:id;`

	result, err := s.db.NamedExec(sqlStatement, &book)

	if err != nil {
		return nil, err
	}

	return result, nil
}
