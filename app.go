package main

import (
	"fmt"
	"nhannkl/demo-golang/utils"
	"nhannkl/demo-golang/view"
)

func main() {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("===== Students & Teachers Management =====")
		fmt.Println("1. Students Management.")
		fmt.Println("2. Teachers Management.")
		fmt.Println("3. Exit.")

		var choice int

		choice = utils.ReadInt("Enter your choice: ")

		switch choice {
		case 1:
			fmt.Print("\033[H\033[2J")
			view.ViewStudentMenu()
		case 2:
			fmt.Print("\033[H\033[2J")
			view.ViewTeacherMenu()
		case 3:
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Please choose an option in the Menu")
		}

	}

}
