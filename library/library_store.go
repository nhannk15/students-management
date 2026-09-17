package library

import (
	"fmt"
	"nhannkl/demo-golang/model"
)

type Library struct {
	Book map[string]model.Book
}

func NewLibrary() *Library {
	return &Library{
		Book: make(map[string]model.Book),
	}
}

func (library *Library) AddBookToStore(id, title, author string) error {
	if _, exists := (*library).Book[id]; exists {
		return fmt.Errorf("book with id %s has existed.", id)
	}

	(*library).Book[id] = model.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}

	return nil
}
