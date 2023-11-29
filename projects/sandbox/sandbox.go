package main

import (
	"fmt"
	display "mysandbox/display"
)

func getName(name string) {
	name = "cloud"
}

type personStruct struct {
	name string
	age int
	sayName func()
}

func createPerson(n string) personStruct {
	return personStruct{
		name: n,
		age: 11,
		sayName: func() {
			fmt.Println(age)
		},
	}
}

func main() {
  display.Display()
	display.Kek()

	phonebook := map[string]int {
		"mario": 345543245,
	}

	fmt.Println(phonebook["mario"])

	x := "tifa"

	getName(x)

	newCharacter := createPerson("Abed")

	fmt.Println(x)
	newCharacter.sayName()
}