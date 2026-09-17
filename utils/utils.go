package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

func PauseProgram() {
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanf("%d")
	ClearConsole()
}

func GenerateId() string {
	return uuid.New().String()
}

func ReadInt(message string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	for {
		tempNumber, err := reader.ReadString('\n')
		tempNumber = strings.TrimSpace(tempNumber)
		if err != nil {
			fmt.Print("Error occured, please enter again: ")
			continue
		}

		realNumber, err := strconv.Atoi(tempNumber)
		if err != nil {
			fmt.Println(err)
			fmt.Print("You must enter an Integer, please enter again: ")
			continue
		}

		return realNumber
	}

}

func ReadString(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	for {
		tempString, err := reader.ReadString('\n')
		tempString = strings.TrimSpace(tempString)
		if err != nil {
			fmt.Print("Error occured, please try again: ")
			continue
		}

		if tempString == "" {
			fmt.Print("Must contain at least one character, please try again: ")
			continue
		}

		return tempString
	}

}
