package main

import (
	"fmt"
	"price-calculator/prices"
)

func main() {
	withoutTaxPrices := []float64{10.0, 20.0, 30.0}
	taxes := []float64{0, 0.07, 0.1, 0.15}

	result := prices.CalculatePricesWithTax(taxes, withoutTaxPrices) 



	fmt.Println(result)
}