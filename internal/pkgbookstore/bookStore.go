package pkgbookstore

import (
	"errors"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
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
func (bs BookStore) Delete(ID int) (error) {
	_, ok := bs.Books[ID]
	if ok {
		delete(bs.Books, ID)
		return nil
	}
	return errors.New("element is missing")
}

// FindByID ... Read
func (bs BookStore) FindByID(ID int) *Book {
	return bs.Books[ID]
}

// UpdateByID ... Update book replacing
func (bs BookStore) UpdateByID(b Book) (error) {
	_, ok := bs.Books[b.ID]
	if ok {
		bs.Books[b.ID] = &b
		return nil
	}
	return errors.New("element is missing")	
}

// CreateFile ...
func CreateFile(filename string,) (*os.File	, error) {
	file, err := os.Create(filename)
	if err != nil {
			return nil, errors.New("file not was created")	
	}

	return file, nil
}

// WriteFile ...
	func WriteFile(file *os.File, b *Book) (int, error) {

	data, _ := json.MarshalIndent(b, "", " ")
	n, err := file.WriteString(string(data) + "\n")
	if err != nil {
		file.Close()
		return 0, errors.New("file not was writing")	
	}
	return n, nil
}

// ReadFile ...
func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("the file cannot be read")	
	}
	return data, nil
}