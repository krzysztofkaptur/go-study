package main

import (
	"fmt"
	"shop/product"
	"shop/user"
)

func main() {
	user1 := user.New("Todd")
	product1 := product.New("Scissors", 2.99)

	user1.BuyProduct(product1)

	fmt.Println(user1.ShowCart())
}