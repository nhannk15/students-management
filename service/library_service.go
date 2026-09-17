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

	fmt.Printf("%s %s %s\n", id, title, author)
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

func AddBorrower() error {
	return nil
}

func ListBorrowers() error {
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
