package main

import (
	"fmt"
)

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type German struct {
	name string
}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(name string) string {
	return fmt.Sprintf("Hallo %v!", name)
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %v: %v", g.LanguageName(), g.Greet(name))
}

func main() {
	germanGreeter := German{}

	fmt.Println(SayHello("Dietrich", germanGreeter))
}