package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getFullName() string {
	fmt.Println("What is your full name? ")

	inputReader := bufio.NewReader(os.Stdin)

	input, _ := inputReader.ReadString('\n')
	
	return input	
}

func getInitials(n string) []string {
	capitalized := strings.ToUpper(n)
	
	names := strings.Split(capitalized, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	return initials
}

func main() {
	fullName := getFullName()

	initials := getInitials(fullName)

	fmt.Print(strings.Join(initials, ""))
}