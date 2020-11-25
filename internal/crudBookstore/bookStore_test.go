package crudbookstore

import (
	"testing"
)

func TestBookStoreAdd(t *testing.T) {
	var bookID = 123

	bookStore := NewBookStore()

	bookFinded := bookStore.FindByID(bookID)
	if bookFinded != nil {
		t.Errorf("El libro con el ID %d ya existe!!!\n", bookID)
	}

	newBook := &Book{
		ID:        123,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1300,
	}
	bookStore.Add(*newBook)
	bookFinded = bookStore.FindByID(bookID)
	if bookFinded == nil {
		t.Errorf("El libro con el ID %d no fue agregado!!!\n", bookID)
	}

	if bookFinded.Name != newBook.Name {
		t.Errorf("El libro con el ID %d no tiene el nombre de la consulta", bookID)
	}
}


func TestBookStoreDelete(t *testing.T) {
	var bookID = 123

	bookStore := NewBookStore()

	errDelete1 := bookStore.Delete(bookID)
	if errDelete1 != nil {
		t.Logf("The book with ID% d cannot be deleted because it does not exist!!!\n", bookID)
	}
	
	newBook := &Book{
		ID:        123,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1300,
	}
	bookStore.Add(*newBook)
	
	errDelete2 := bookStore.Delete(bookID)
	if errDelete2 == nil {
		t.Logf("Delete with id %v was succesfully!!!\n", bookID)
	}
}
