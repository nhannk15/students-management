package library

import (
	"fmt"
	"nhannkl/demo-golang/model"
)

type Library struct {
	Book      map[string]model.Book
	Borrowers map[string]model.Borrower
}

func NewLibrary() *Library {
	return &Library{
		Book:      make(map[string]model.Book),
		Borrowers: make(map[string]model.Borrower),
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
