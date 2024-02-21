package user

import "fmt"

type User struct {
	name string
}

func New(n string) User {
	return User{
		name: n,
	}
}

func (u User) SayName() string {
	return u.name
}

func (u User) Greet() string {
	return fmt.Sprintf("Hi my name is %v", u.name)
}

func (u User) Dupa() string {
	return "dupa"
}