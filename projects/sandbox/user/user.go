package user

import (
	"fmt"
)

type user struct {
	name string
}

func CreateUser(name string) user {
	user := user{
		name: name,
	}

	return user
}

func User() {
	fmt.Print("User init")
}