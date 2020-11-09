package main

import (
	"fmt"
	"tudai.seminario.golang.practica/internal/pkgbookstore"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	newBook := &pkgbookstore.Book{
		ID:        123,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1300,
	}

	newBook2 := &pkgbookstore.Book{
		ID:        456,
		Name:      "Salem`s Lot",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1500,
	}

	// Porque al newBook2 tengo que ponerle * para generar una copia?
	// Despues cuando el newBook se lo paso al find, tengo que pasalo sin el * a diferencia del newBook2
	// Porque esa diferencia? Como seria generar una copia para pasarlo como parametro de la misma forma?
	newBook3 := *newBook2
	newBook3.ID = 678

	var findBookID = 456
	var deleteBookID = 456

	fmt.Println("\n------------------------ CREATE ------------------------")
	bookStore := pkgbookstore.NewBookStore()
	bookStore.Add(*newBook)
	bookStore.Add(*newBook2)
	bookStore.Print()

	fmt.Println("\n------------------------ READ --------------------------")
	bookFinded := bookStore.FindByID(findBookID)
	if bookFinded != nil {
		fmt.Printf("The book with id %v was found successfully!!!\n", findBookID)
	} else {
		fmt.Printf("The book with id %v was not!!!\n", findBookID)
	}

	fmt.Println("\n------------------------ UPDATE ------------------------")
	newBook2.Price = 3999
	errUpdate := bookStore.UpdateByID(*newBook2)
	if errUpdate != nil {
		fmt.Printf("Update Book with id %v: %v\n", newBook2.ID, errUpdate)
	} else {
		fmt.Printf("Update Book with id %v was succesfully\n", newBook2.ID)
	}
	// errUpdate := bookStore.UpdateByID(newBook3)
	// if errUpdate != nil {
	// 	fmt.Printf("Update Book with id %v: %v\n", newBook3.ID, errUpdate)
	// } else {
	// 	fmt.Printf("Update Book with id %v was succesfully\n", newBook3.ID)
	// }
	bookStore.Print()

	fmt.Println("\n------------------------ DELETE ------------------------")
	errDelete := bookStore.Delete(deleteBookID)
	if errDelete != nil {
		fmt.Printf("Delete with id %v: %v\n", deleteBookID, errDelete)
	} else {
		fmt.Printf("Delete with id %v was succesfully\n", deleteBookID)
	}
	bookStore.Print()
	fmt.Println("")
	
	/*******************************************************************************/

	newBook4 := &pkgbookstore.Book{
		ID:        1,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1300,
	}

	newBook5 := &pkgbookstore.Book{
		ID:        2,
		Name:      "Salem`s Lot",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1500,
	}

	newBook6 := &pkgbookstore.Book{
		ID:        3,
		Name:      "Salem`s Lot",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Price:     1500,
	}

	bookStore2 := pkgbookstore.NewBookStore()
	bookStore2.Add(*newBook4)
	bookStore2.Add(*newBook5)
	bookStore2.Add(*newBook6)

	var filename = "files/test.txt"

	file, err := pkgbookstore.CreateFile(filename);
	if err != nil {
		fmt.Println(err)
	}

	for _, b := range bookStore2.Books {
		l, err := pkgbookstore.WriteFile(file, b)
		fmt.Println(l, "bytes written successfully")
		if err != nil {
			fmt.Println(err)
		}
	}

	defer file.Close()

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := pkgbookstore.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
