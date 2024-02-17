package main

import (
	"fmt"

	concurrencypractice "github.com/leonardodelira/100-go-mistakes/8-concurrency-practice"
)

func main() {
	customer := concurrencypractice.Customer{}
	err := customer.UpdateAge(-21)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(customer.ShowAge())
}
