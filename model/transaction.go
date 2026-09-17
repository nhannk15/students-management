package model

import "time"

type Transaction struct {
	Id            string
	BookId        string
	BorrowerEmail string
	BorrowDate    time.Time
	ReturnDate    time.Time
}
