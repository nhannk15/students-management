package library

import (
	"fmt"
	"nhannkl/demo-golang/model"
	"time"
)

type Library struct {
	Book         map[string]model.Book
	Borrowers    map[string]model.Borrower
	Transactions map[string]model.Transaction
}

func NewLibrary() *Library {
	return &Library{
		Book:         make(map[string]model.Book),
		Borrowers:    make(map[string]model.Borrower),
		Transactions: make(map[string]model.Transaction),
	}
}

func (library *Library) AddBookToStore(id, title, author string) error {
	if _, exists := (*library).Book[id]; exists {
		return fmt.Errorf("book with id %s has existed.", id)
	}

	(*library).Book[id] = model.Book{
		Id:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}

	return nil
}

func (library *Library) ListBooksStore() map[string]model.Book {
	return (*library).Book
}

func (library *Library) AddBorrowerToStore(id, name, email string) error {
	if _, exists := (*library).Borrowers[email]; exists {
		return fmt.Errorf("Duplicated email.")
	}

	(*library).Borrowers[email] = model.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return nil
}

func (library *Library) ListBorrowersStore() []model.Borrower {
	borrowers := make([]model.Borrower, 0, len((*library).Borrowers))
	for _, value := range (*library).Borrowers {
		borrowers = append(borrowers, value)
	}

	return borrowers
}

func (library *Library) BorrowBookStore(transactionId, borrowerEmail, bookId string) error {

	book, bookExists := (*library).Book[bookId]
	if !bookExists {
		return fmt.Errorf("Book didn't exist")
	}

	if book.IsBorrowed {
		return fmt.Errorf("Book has been borrowed")
	}

	_, borrowerExists := (*library).Borrowers[borrowerEmail]
	if !borrowerExists {
		return fmt.Errorf("Borrower didn't exist")
	}

	//--- Reassign
	book.IsBorrowed = true
	(*library).Book[bookId] = book

	(*library).Transactions[transactionId] = model.Transaction{
		Id:            transactionId,
		BorrowerEmail: borrowerEmail,
		BookId:        bookId,
		BorrowDate:    time.Now(),
	}

	fmt.Printf("%+v\n", (*library).Transactions[transactionId])

	return nil
}
