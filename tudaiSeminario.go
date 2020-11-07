package main

import (
	"fmt"
)

// Book ...
type Book struct {
	ID        int
	name      string
	language  string
	status    string
	genre     string
	editorial string
	author    string
	price     int
}

// BooksStore ...
type BooksStore struct {
	books map[int]*Book
}

// NewBooksStore ...
func NewBooksStore() BooksStore {
	books := make(map[int]*Book)
	return BooksStore{
		books,
	}
}

// Add ... Create
func (bs BooksStore) Add(b Book) {
	bs.books[b.ID] = &b
}

// Print ...
func (bs BooksStore) Print() {
	for k, v := range bs.books {
		// fmt.Printf("%v\t[%v]\t%v\n", k, v.ID, v.name)
		// fmt.Printf("%#v\n", *v)
		fmt.Printf("%v\t%+v\n", k, *v)
		// jsonF, _ := json.Marshal(v)
		// fmt.Printf(string(jsonF))
	}
}

// Delete ... Delete
func (bs BooksStore) Delete(ID int) {
	delete(bs.books, ID)
}

// FindByID ... Read
func (bs BooksStore) FindByID(ID int) *Book {
	return bs.books[ID]
}

// UpdateByID ... Update book replacing
func (bs BooksStore) UpdateByID(b Book) {
	bs.books[b.ID] = &b
}

func main() {

	booksStore := NewBooksStore()

	newBook := &Book{
		ID:        123,
		name:      "It",
		language:  "Spanish",
		status:    "New",
		genre:     "Terror",
		editorial: "Plaza&James",
		author:    "Stephen King",
		price:     1300,
	}

	newBook2 := &Book{
		ID:        456,
		name:      "Salem`s Lot",
		language:  "Spanish",
		status:    "New",
		genre:     "Terror",
		editorial: "Plaza&James",
		author:    "Stephen King",
		price:     1500,
	}

	booksStore.Add(*newBook)
	booksStore.Add(*newBook2)
	newBook2.price = 3999
	bookFinded := booksStore.FindByID(456)
	if bookFinded != nil {
		fmt.Printf("Finded!!!")
	}
	booksStore.UpdateByID(*newBook2)
	booksStore.Delete(456)

	booksStore.Print()

}
