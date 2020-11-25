package crudbookstore

import (
	"testing"
)

func TestBookStoreAdd(t *testing.T) {
	var bookID = 1

	newBook := &Book{
		ID:        1,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1300,
	}

	bookStore := NewBookStore()

	bookFinded := bookStore.FindByID(bookID)
	if bookFinded != nil {
		t.Errorf("El libro con el ID %d ya existe!!!\n", bookID)
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

func TestBookStoreRead(t *testing.T) {
	var bookID = 1

	newBook := &Book{
		ID:        1,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1300,
	}

	bookStore := NewBookStore()

	bookStore.Add(*newBook)
	bookFinded := bookStore.FindByID(bookID)
	if bookFinded == nil {
		t.Errorf("El libro con el ID %d no fue encontrado en la bookstore!!!\n", bookID)
	}
}

func TestBookStoreUpdate(t *testing.T) {
	var bookID = 1

	newBook := &Book{
		ID:        1,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1300,
	}

	bookStore := NewBookStore()
	bookStore.Add(*newBook)

	newBook.Price = 1500

	bookStore.Update(*newBook)
	bookFinded := bookStore.FindByID(bookID)
	if bookFinded.Price != 1500 {
		t.Errorf("El libro con el ID %d no se pudo eliminar!!!\n", bookID)
	}
}

func TestBookStoreDelete(t *testing.T) {
	var bookID = 1

	newBook := &Book{
		ID:        1,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1300,
	}

	bookStore := NewBookStore()

	bookStore.Add(*newBook)

	bookStore.Delete(bookID)

	bookFinded := bookStore.FindByID(bookID)
	if bookFinded != nil {
		t.Errorf("El libro con el ID %d no se pudo eliminar!!!\n", bookID)
	}
}
