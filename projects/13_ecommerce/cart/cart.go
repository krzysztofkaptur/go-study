package cart

import (
	"shop/product"
)

type Cart struct {
	products []product.Product
}

func New() Cart {
	return Cart{
		products: []product.Product{},
	}
}

func (c *Cart) AddProduct(p product.Product) {
	c.products = append(c.products, p)
}