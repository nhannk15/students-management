package main

import (
	"fmt"
	"nhannkl/demo-golang/service"
	"nhannkl/demo-golang/utils"
)

func main() {
	utils.ClearConsole()
	for {
		fmt.Println("===== Library Management System =====")
		fmt.Println("1. Add book.")
		fmt.Println("2. List books.")
		fmt.Println("3. Add borrower.")
		fmt.Println("4. List all borrowers.")
		fmt.Println("5. Borrow book.")
		fmt.Println("6. Check borrowing history.")
		fmt.Println("7. Return book.")
		fmt.Println("8. Search by title or author.")
		fmt.Println("9. Exit.")
		choice := utils.ReadInt("---> Enter your choice: ")

		switch choice {
		case 1:
			if err := service.AddBook(); err != nil {
				fmt.Printf("Error occured when adding book: %v \n", err)
			}
		case 2:
			if err := service.ListBooks(); err != nil {
				fmt.Printf("Error occured when listing books: %v \n", err)
			}
		case 3:
			if err := service.AddBorrower(); err != nil {
				fmt.Printf("Error occured when adding borrower: %v \n", err)
			}
		case 4:
			if err := service.ListBorrowers(); err != nil {
				fmt.Printf("Error occured when listing borrowers: %v \n", err)
			}
		case 5:
			if err := service.BorrowBook(); err != nil {
				fmt.Printf("Error occured when borrowing book: %v \n", err)
			}
		case 6:
			if err := service.ListBorrowHistory(); err != nil {
				fmt.Printf("Error occured when listing borrowing history: %v \n", err)
			}
		case 7:
			if err := service.ReturnBook(); err != nil {
				fmt.Printf("Error occured when returning book: %v \n", err)
			}
		case 8:
			if err := service.SearchBooks(); err != nil {
				fmt.Printf("Error occured when searching books: %v \n", err)
			}
		case 9:
			return
		}

		utils.PauseProgram()
	}
}
