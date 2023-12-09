package main

import (
	"fmt"
	"mysandbox/bill"
)

type shape interface {
	area() float64
}

type shapeStruct struct {
	x float64
	y float64
}

func (s shapeStruct) area() float64 {
	return s.x * s.y
}

func main() {
	myBill := bill.NewBill("Mario's bill")


	marioBill := map[string]float64{
		"salami": 4.47,
		"salad": 7.47,
		"steak": 12.47,
	}

	myBill.AddBill(marioBill)
	myBill.AddTip(5.0)
	myBill.AddItem("french fries", 4.34)
	fmt.Println(myBill.Format())

	myShape := shapeStruct{
		x: 5.5,
		y: 15,
	}

	fmt.Println(myShape.area())
}