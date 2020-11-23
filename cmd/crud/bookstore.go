package main

import (
	"fmt"

	"tudai.seminario.golang.practica/internal/crudbookstore"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	newBook := &crudbookstore.Book{
		ID:        1,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Publicado: "01-07-1986",
		Price:     1300.99,
	}

	newBook2 := &crudbookstore.Book{
		ID:        2,
		Name:      "Salem`s Lot",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Publicado: "17-10-1975",
		Price:     1500,
	}

	// Porque al newBook2 tengo que ponerle * para generar una copia?
	// Despues cuando el newBook se lo paso al find, tengo que pasalo sin el * a diferencia del newBook2
	// Porque esa diferencia? Como seria generar una copia para pasarlo como parametro de la misma forma?
	newBook3 := *newBook2
	newBook3.ID = 3

	var findBookID = 2
	var deleteBookID = 2

	/*******************************************************************************/
	/********************************** CRUD ***************************************/
	/*******************************************************************************/

	fmt.Println("\n------------------------ CREATE ------------------------")
	bookStore := crudbookstore.NewBookStore()
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
	/**************************** Write/Read File **********************************/
	/*******************************************************************************/

	newBook4 := &crudbookstore.Book{
		ID:        1,
		Name:      "It",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Publicado: "01-07-1986",
		Price:     1300.99,
	}

	newBook5 := &crudbookstore.Book{
		ID:        2,
		Name:      "Salem`s Lot",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: "Plaza&James",
		Author:    "Stephen King",
		Publicado: "17-10-1975",
		Price:     1500,
	}

	newBook6 := &crudbookstore.Book{
		ID:        3,
		Name:      "Carrie",
		Language:  "Spanish",
		Status:    "New",
		Genre:     "Terror",
		Editorial: `Plaza&James`,
		Author:    "Stephen King",
		Publicado: "05-04-1974",
		Price:     1750.50,
	}

	bookStore2 := crudbookstore.NewBookStore()
	bookStore2.Add(*newBook4)
	bookStore2.Add(*newBook5)
	bookStore2.Add(*newBook6)

	var filename = "files/test.txt"

	file, err := crudbookstore.CreateFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	for _, b := range bookStore2.Books {
		l, err := crudbookstore.WriteFile(file, b)
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

	data, err := crudbookstore.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
