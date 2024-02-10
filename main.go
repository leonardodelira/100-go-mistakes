package main

import (
	"fmt"

	errormanagement "github.com/leonardodelira/100-go-mistakes/6-error-management"
)

func main() {
	err := errormanagement.AddContextError()
	if err != nil {
		fmt.Print(err.Error())
	}
}
