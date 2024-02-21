package user

import (
	"shop/cart"
	"shop/product"
	"time"
)

type User struct {
	name      string
	createdAt time.Time
	updatedAt time.Time
	cart cart.Cart
}

func New(name string) User {
	newCart := cart.New()

	return User{
		name: name,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		cart: newCart,
	}
}

func (u *User) BuyProduct(p product.Product) {
	u.cart.AddProduct(p)
}

func (u *User) ShowCart() cart.Cart {
	return u.cart
}

func (u *User) ShowName() string {
	return u.name
}