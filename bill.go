package main

import "fmt"

type bill struct {
	name  string
	items map[string]float32
	tip   int
}

// make bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float32{"india":23.4,"america":121.43,}, tip: 0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill\n"
	var total float32 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":",v)
		total+=v
	}
	fs+=fmt.Sprintf("%-25v ...$%0.2f","total:",total)
	return fs
}