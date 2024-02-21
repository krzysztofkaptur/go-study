package product

type Product struct {
	name  string
	price float64
}

func New(name string, price float64) Product {
	return Product{
		name:  name,
		price: price,
	}
}