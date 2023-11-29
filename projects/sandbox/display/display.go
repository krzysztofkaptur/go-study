package display

import "fmt"

func changeName(n *string) {
	*n = "Pierce"
}

func Display() {
	x := "Todd"

	y := &x

	x = "Jeff"
	
	*y = "Annie"

	changeName(y)

	fmt.Println(x, *y)
}