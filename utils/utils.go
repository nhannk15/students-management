package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadInt(message string) int {
	fmt.Printf("%s", message)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error occured...")
			continue
		}
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("You must enter an integer")
			continue
		}
		return number

	}
}

func ReadString(message string) string {
	fmt.Printf("%s", message)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error occured...")
			continue
		}
		input = strings.TrimSpace(input)
		return input
	}
}

func PauseProgram() {
	fmt.Printf("Press 'Enter' to continue...")
	fmt.Scanf("%s")
}
