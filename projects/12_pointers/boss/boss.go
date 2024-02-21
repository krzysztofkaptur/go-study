package boss

import (
	"fmt"
)

type Boss struct {
	name string
}

func New(n string) Boss {
	return Boss{
		name: n,
	}
}

func (b Boss) SayName() string {
	return b.name
}

func (b Boss) Greet() string {
	return fmt.Sprintf("Hi my name is %v", b.name)
}