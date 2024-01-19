package main

import controlstructs "github.com/leonardodelira/100-go-mistakes/3-control-structs"

func main() {
	customers := []controlstructs.Customer{
		{
			ID:      1,
			Balance: 100,
		},
		{
			ID:      2,
			Balance: 0,
		},
		{
			ID:      3,
			Balance: -30,
		},
	}
	store := controlstructs.Store{}
	store.ExampleCorrect32(customers)
}
