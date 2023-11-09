package datatypes

import "fmt"

/*
Mistake 23: Not properly checking if a slice is empty
*/

//Em resumo é uma boa pratica verificar se o slice é nil ou empty por fazer uso da função len()

func ExampleMistake23() {
	result := OtherFunc("")
	result2 := OtherFunc("asd")

	fmt.Print(len(result) != 0)  //false
	fmt.Print(len(result2) != 0) //false

	//No cenário acima, usar o len() é útil mesmo se o slice estiver nil ou empty
}

func OtherFunc(d string) []float64 {
	var operations = make([]float64, 0)
	if d == "" {
		return nil
	}
	return operations
}
