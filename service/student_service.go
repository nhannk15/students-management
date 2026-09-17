package service

import (
	"fmt"
	"nhannkl/demo-golang/model/student"
	"nhannkl/demo-golang/utils"
	"slices"
	"strconv"
)

var studentList []student.Student

func AddStudent() {
	var id = utils.ReadInt("Enter student's id: ")
	var name = utils.ReadString("Enter student's name: ")
	var className = utils.ReadString("Enter class name: ")
	var numOfGrades = utils.ReadInt("Enter the number of grades: ")
	var listOfGrades []int
	for i := 0; i < numOfGrades; i++ {
		var grade = utils.ReadInt("Enter grade " + strconv.Itoa(i+1) + ": ")
		listOfGrades = append(listOfGrades, grade)
	}

	newStudent := student.Student{
		Id:        id,
		Name:      name,
		ClassName: className,
		Grades:    listOfGrades,
	}

	studentList = append(studentList, newStudent)
	fmt.Print("Student added. ")

	utils.PauseProgram()
}

func ListStudent() {
	if len(studentList) == 0 {
		fmt.Println("Empty student list")
		utils.PauseProgram()
		return
	}
	fmt.Printf("%-10v | %-10v | %-10v | %-10v \n",
		"Id",
		"Name",
		"Class Name",
		"Average",
	)

	for _, student := range studentList {
		fmt.Printf("%-10v | %-10v | %-10v | %-10v \n",
			student.Id,
			student.Name,
			student.ClassName,
			student.CalculateAverageGrades(),
		)
	}

	utils.PauseProgram()
}

func UpdateStudent() {
	for {
		var id int = utils.ReadInt("Enter the student id you want to update: ")
		var index int = slices.IndexFunc(studentList, func(student student.Student) bool {
			return student.Id == id
		})
		if index == -1 {
			fmt.Println("The student you want does not exist")
		} else {
			var student *student.Student = &studentList[index]
			var temp string
			fmt.Println("Press 'Enter' for keeping the curent data...")
			temp = utils.ReadString("Enter name (" + student.Name + "): ")
			if temp != "" {
				(*student).Name = temp
			} else {
				fmt.Println("Kept...")
			}

			temp = utils.ReadString("Enter class (" + student.ClassName + "): ")
			if temp != "" {
				(*student).ClassName = temp
			} else {
				fmt.Println("Kept...")
			}

			for i := 0; i < len(student.Grades); i++ {
				fmt.Printf("Update Grade %d: (%d): ", i+1, student.Grades[i])
				for {
					var tempNumber = utils.ReadString("")
					if tempNumber == "" {
						fmt.Println("Kept...")
						break
					}
					finalNumber, err := strconv.Atoi(tempNumber)
					if err != nil {
						fmt.Println("You must enter an integer!")
						continue
					} else {
						(*student).Grades[i] = finalNumber
						break
					}
				}

			}
			break
		}
	}
	fmt.Print("Student updated. ")
	utils.PauseProgram()
}

func DeleteStudent() {

	var id int = utils.ReadInt("Enter the student id you want to update: ")
	var index int = slices.IndexFunc(studentList, func(student student.Student) bool {
		return student.Id == id
	})
	if index == -1 {
		fmt.Println("The student you want does not exist")
	} else {
		studentList = slices.DeleteFunc(studentList, func(student student.Student) bool {
			return student.Id == id
		})
		fmt.Print("Student deleted. ")
	}
	utils.PauseProgram()

}

func FindStudent() {
	var id int = utils.ReadInt("Enter the student id you want to update: ")
	var index int = slices.IndexFunc(studentList, func(student student.Student) bool {
		return student.Id == id
	})
	if index == -1 {
		fmt.Println("The student you want does not exist")
	} else {
		var student = studentList[index]
		fmt.Printf("%-10v | %-10v | %-10v | %-10v \n",
			"Id",
			"Name",
			"Class Name",
			"Average",
		)
		fmt.Printf("%-10v | %-10v | %-10v | %-10v \n",
			student.Id,
			student.Name,
			student.ClassName,
			student.CalculateAverageGrades(),
		)
		fmt.Print("Student found. ")
	}
	utils.PauseProgram()
}
