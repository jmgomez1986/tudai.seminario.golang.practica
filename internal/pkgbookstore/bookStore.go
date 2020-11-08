package pkgbookstore

import (
	"fmt"
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
	Price     int
}

// BookStore ...
type BookStore struct {
	Books map[int]*Book
}

// NewBookStore ...
func NewBookStore() BookStore {
	booksMap := make(map[int]*Book)
	return BookStore{
		Books: booksMap,
	}
}

// Add ...
func (bs BookStore) Add(b Book) {
	bs.Books[b.ID] = &b
}

// Print ...
func (bs BookStore) Print() {
	for k, v := range bs.Books {
		// fmt.Printf("%v\t[%v]\t%v\n", k, v.ID, v.name)
		fmt.Printf("%v\t%+v\n", k, *v)
	}
}

// Delete ... 
func (bs BookStore) Delete(ID int) {
	delete(bs.Books, ID)
}

// FindByID ... Read
func (bs BookStore) FindByID(ID int) *Book {
	return bs.Books[ID]
}

// UpdateByID ... Update book replacing
func (bs BookStore) UpdateByID(b Book) {
	bs.Books[b.ID] = &b
}