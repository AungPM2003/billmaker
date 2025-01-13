package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}
func (b *bill) addingItems(item string, price float64) {
	b.items[item] = price
}
func (b *bill) updatingTip(tip float64) {
	b.tip = tip
}
func (b *bill) format() string {
	fs := fmt.Sprintf("%12v\n", b.name)
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf(" %-10v ...%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf(" %-10v ...%0.2f\n", "Tip: ", b.tip)
	fs += fmt.Sprintf(" %-10v ...%0.2f", "Total:", total+b.tip)
	return fs
}
