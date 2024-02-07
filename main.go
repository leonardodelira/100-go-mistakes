package main

import (
	"log"

	functionsandmethods "github.com/leonardodelira/100-go-mistakes/5-functions-and-methods"
)

func main() {
	customer := &functionsandmethods.Customer{
		Name: "Test",
		Age:  1,
	}
	if err := customer.Validate(); err != nil {
		log.Fatalf("customer is not valid %v", err)
	}
}
