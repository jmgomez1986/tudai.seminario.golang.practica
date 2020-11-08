package main

import (
	"fmt"
)

func main() {

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

	newBook2 := &Book{
		ID:        456,
		Name:      "Salem`s Lot",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1500,
	}

	var bookID = 456

	// Create
	bookStore := NewBookStore()
	bookStore.Add(*newBook)
	bookStore.Add(*newBook2)
	fmt.Println("\n------------------------ CREATE ------------------------")
	bookStore.Print()
	// Read
	bookFinded := bookStore.FindByID(bookID)
	fmt.Println("\n------------------------ READ --------------------------")
	if bookFinded != nil {
		fmt.Printf("The book with id %v was found successfully!!!\n", bookID)
	}
	// Update
	newBook2.Price = 3999
	bookStore.UpdateByID(*newBook2)
	fmt.Println("\n------------------------ UPDATE ------------------------")
	bookStore.Print()
	// Delete
	bookStore.Delete(bookID)
	fmt.Println("\n------------------------ DELETE ------------------------")
	bookStore.Print()
}
