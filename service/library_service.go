package service

import (
	"fmt"
	"nhannkl/demo-golang/library"
	"nhannkl/demo-golang/utils"
)

func AddBook(library *library.Library) error {
	id := utils.GenerateId()
	title := utils.ReadString("Enter title: ")
	author := utils.ReadString("Enter author: ")

	if err := library.AddBookToStore(id, title, author); err != nil {
		return err
	}

	fmt.Println("Successfully added a new Book")
	return nil
}

func ListBooks(library *library.Library) error {
	books := (*library).ListBooksStore()
	if len(books) == 0 {
		fmt.Println("Empty library!")
		return nil
	}

	for _, value := range (*library).Book {
		var bookStatus string = "Available"
		if value.IsBorrowed {
			bookStatus = "Borrowed"
		}
		fmt.Printf("Id: %s - Title: %s - Author: %s - %s\n", value.Id, value.Title, value.Author, bookStatus)
	}
	return nil
}

func AddBorrower(library *library.Library) error {
	id := utils.GenerateId()
	name := utils.ReadString("Enter name: ")
	email := utils.ReadString("Enter email: ")

	if err := (*library).AddBorrowerToStore(id, name, email); err != nil {
		return err
	}

	fmt.Println("Successfully added a new Borrower")

	return nil
}

func ListBorrowers(library *library.Library) error {
	borrowers := (*library).ListBorrowersStore()
	if len(borrowers) == 0 {
		fmt.Println("Empty borrower list.")
		return nil
	}

	for _, value := range borrowers {
		fmt.Printf("Id: %s - Name: %s - Email: %s\n", value.Id, value.Name, value.Email)
	}
	return nil
}

func BorrowBook() error {
	return nil
}

func ListBorrowHistory() error {
	return nil
}

func ReturnBook() error {
	return nil
}

func SearchBooks() error {
	return nil
}
