package main

import (
	"fmt"
)

func filter(list []int, transform func(num int) bool) []int {
	result := []int{}

	for i := 0; i < len(list); i++ {
		if transform(list[i]) {
			result = append(result, list[i])
		}
	}
	
	return result
}

func moreThan4(num int) bool {
	return num > 4
}

func oddNumbers(num int) bool {
	return num % 2 != 0
}

func main() {
	list := []int{1,2,3,4,5,6,7,8,9}

	filteredList := filter(list, oddNumbers)

	fmt.Println(filteredList)
}