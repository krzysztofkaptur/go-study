package main

import (
	"fmt"
	"pointer/boss"
	"pointer/user"
	"pointer/values"
)

type Namer interface {
	SayName() string
	Greet() string
}

func main() {
	playWithUser()
	playWithBoss()
}

func playWithUser() {
	user1 := user.New("Todd")
	
	// fmt.Println(user1.Greet())
	
	fmt.Println(showName(user1))
}

func playWithBoss() {
	boss1 := boss.New("Jeff")

	// fmt.Println(boss1.Greet())

	fmt.Println(showName(boss1))
}


func showName(person Namer) string {
	return person.Greet()
}


func playWithNumbers() {
	someVal := 0
	
	values.Increment(&someVal)
	
	fmt.Println(someVal)
}