package service

import (
	"fmt"
	"nhannkl/demo-golang/model/student"
	"nhannkl/demo-golang/utils"
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
	fmt.Println("Student added")

	utils.PauseProgram()
}

func ListStudent() {
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
