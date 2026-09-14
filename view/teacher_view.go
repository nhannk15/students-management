package view

import (
	"fmt"
	"nhannkl/demo-golang/utils"
)

func ViewTeacherMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("===== Teachers Management =====")
	fmt.Println("1. Add student.")
	fmt.Println("2. Delete Student.")
	fmt.Println("3. Update student.")
	fmt.Println("4. List student.")
	fmt.Println("5. Find student.")
	fmt.Println("6. Return.")

	var choice int
	choice = utils.ReadInt("Enter your choice: ")
	fmt.Scanf("%d")

	switch choice {
	case 1:
		fmt.Print("\033[H\033[2J")
		fmt.Println("Student added")
	case 2:
		fmt.Print("\033[H\033[2J")
		fmt.Println("Student added")
	case 3:
		fmt.Print("\033[H\033[2J")
		fmt.Println("Student added")
	case 4:
		fmt.Print("\033[H\033[2J")
		fmt.Println("Student added")
	case 5:
	case 6:
		return
	default:
		fmt.Print("\033[H\033[2J")
		fmt.Println("Please choose an option in the Menu")
	}
}
