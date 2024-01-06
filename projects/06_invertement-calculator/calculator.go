package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		inflationRate = 2.5
	)

	invertmentAmount := 1000.0
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Println("Give me an investment amount")
	fmt.Scan(&invertmentAmount)

	fmt.Println("Give me years")
	fmt.Scan(&years)

	fmt.Println("Give me rate")
	fmt.Scan(&expectedReturnRate)
	
	futureValue := invertmentAmount * math.Pow(1 + expectedReturnRate / 100, years)  
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue, futureRealValue)
}