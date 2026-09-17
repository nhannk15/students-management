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

func BorrowBook(library *library.Library) error {
	transactionId := utils.GenerateId()
	borrowerEmail := utils.ReadString("Enter borrower email: ")
	bookId := utils.ReadString("Enter book id: ")

	if err := (*library).BorrowBookStore(transactionId, borrowerEmail, bookId); err != nil {
		return err
	}

	fmt.Println("Borrower successfully borrowed a book.")
	return nil
}

func ListBorrowHistory(library *library.Library) error {
	borrowerEmail := utils.ReadString("Enter borrower email: ")
	transactions := (*library).BorrowHistoryStore(borrowerEmail)
	if len(transactions) == 0 {
		fmt.Println("Borrower hasn't triggered any transaction.")
	} else {
		for _, transaction := range transactions {
			returnStatus := "Not yet"
			if !transaction.ReturnDate.IsZero() {
				returnStatus = transaction.ReturnDate.Format("02/01/2006")
			}
			fmt.Printf("TransactionId: %s - Book: %s - Date Borrowed: %v - Date Returned: %v\n",
				transaction.Id, transaction.BookId, transaction.BorrowDate.Format("02/01/2006"), returnStatus)
		}
	}
	return nil
}

func ReturnBook(library *library.Library) error {
	transactionId := utils.ReadString("Enter Transaction Id: ")
	if err := (*library).ReturnBookStore(transactionId); err != nil {
		return err
	}

	fmt.Println("Successfully returned a book")
	return nil
}

func SearchBooks(library *library.Library) error {
	searchValue := utils.ReadString("Enter search value: ")
	books := (*library).SearchBookStore(searchValue)

	for _, value := range books {
		var bookStatus string = "Available"
		if value.IsBorrowed {
			bookStatus = "Borrowed"
		}
		fmt.Printf("Id: %s - Title: %s - Author: %s - %s\n", value.Id, value.Title, value.Author, bookStatus)
	}
	return nil
}
