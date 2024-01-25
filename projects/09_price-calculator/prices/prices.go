package prices

func CalculatePricesWithTax(taxes []float64, prices []float64) map[float64][]float64 {
	result := map[float64][]float64{}

	for _, tax := range taxes {
		taxIncludedPrices := make([]float64, len(prices))

		for priceIdx, price := range prices {
			taxIncludedPrices[priceIdx] = price * (1 + tax)
		}

		result[tax] = taxIncludedPrices
	}

	return result
}