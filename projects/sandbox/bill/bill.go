package bill

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip float64
}

func NewBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip: 0,
	}
}

func (b *bill) Format() string {
	fmt.Println(b.items)
	fs := "Bill breakdown: \n"
	total := 0.0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k + ":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "tip:", b.tip)

	return fs
}

func (b *bill) AddItem(name string, value float64) {
	b.items[name] = value
}

func (b *bill) AddTip(t float64) {
	b.tip = t
}

func (b *bill) AddBill(items map[string]float64) {
	b.items = map[string]float64{
		"salami": 4.47,
		"salad": 7.47,
		"steak": 12.47,
	}
}