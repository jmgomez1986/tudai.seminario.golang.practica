package crudbookstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	Price     float32
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
		fmt.Printf("%v\t%+v\n", k, *v)
	}
}

// Delete ...
func (bs BookStore) Delete(ID int) error {
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

// Update ... Update book replacing
func (bs BookStore) Update(b Book) error {
	_, ok := bs.Books[b.ID]
	if ok {
		bs.Books[b.ID] = &b
		return nil
	}
	return errors.New("element is missing")
}

// CreateFile ...
func CreateFile(path string, filename string) (*os.File, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	file, err := os.Create(filename)
	if err != nil {
		return nil, errors.New("file not was created")
	}

	return file, nil
}

// WriteFile ...
func WriteFile(file *os.File, b *Book) (int, error) {

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", " ")
	jsonEncoder.Encode(b)

	n, err := file.WriteString(bf.String() + "\n")

	if err != nil {
		file.Close()
		return 0, errors.New("file not was writing")
	}
	return n, nil
}

// ReadFile ...
func ReadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("the file cannot be read")
	}
	return string(data), nil
}
