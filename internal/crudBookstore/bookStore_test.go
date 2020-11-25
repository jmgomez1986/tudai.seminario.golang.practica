package crudbookstore

import (
	"testing"
)

func TestBookStoreAdd(t *testing.T) {
	var bookID = 1

	bookStore := NewBookStore()

	bookFinded := bookStore.FindByID(bookID)
	if bookFinded != nil {
		t.Errorf("El libro con el ID %d ya existe!!!\n", bookID)
	}

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
	var bookID = 1

	bookStore := NewBookStore()

	// errDelete1 := bookStore.Delete(bookID)
	// if errDelete1 != nil {
	// 	t.Errorf("El libro con el ID %d no puede eliminarse porque no existe!!!\n", bookID)
	// }

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
	bookStore.Add(*newBook)

	errDelete2 := bookStore.Delete(bookID)
	// if errDelete2 == nil {
	// 	t.Errorf("El libro con el ID %d se elimino correctamente!!!\n", bookID)
	// }
	if errDelete2 != nil {
		t.Errorf("El libro con el ID %d no se pudo eliminar!!!\n", bookID)
	}
}

func TestBookStoreUpdate(t *testing.T) {
	// var bookID = 1

	bookStore := NewBookStore()

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
	bookStore.Add(*newBook)

	newBook3 := *newBook
	newBook3.Price = 1500

	errUpdate2 := bookStore.UpdateByID(newBook3)
	if errUpdate2 != nil {
		t.Errorf("El libro con el ID %d no se pudo actualizar!!!\n", newBook3.ID)
	}
}
