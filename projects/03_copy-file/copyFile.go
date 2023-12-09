package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Herro")

	data, _ := os.ReadFile("file.txt")

	fmt.Println(string(data))
}