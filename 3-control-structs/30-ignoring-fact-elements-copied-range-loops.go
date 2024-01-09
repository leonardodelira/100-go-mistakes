package controlstructs

import "fmt"

/*
Mistake 30: Ignoring the fact that elements are copied in range loops

Devemos ter em mente que em Go tudo que nós asignamos é uma cópia, e não o valor original.
Então dependendo da forma que nós interamos de um loop, podemos estar apenas alterando o valor da cópia ao invez da struct original
*/

type Account struct {
	balance float64
}

// Example
func NotCorrectIncrementValue() {
	accounts := []Account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	for _, v := range accounts {
		v.balance += 1000
	}

	//O resultado desse slice não irá ser alterado, porque em nosso range nós alteramos os valores da cópia e não do slice original.
	fmt.Print(accounts) //resultado: {100, 200, 300}
}

// Se quisermos alterar o valor da struct original, devemos alterar o valor usando o index
func CorrectIncrementeValue() {
	accounts := []Account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	for i := range accounts {
		accounts[i].balance += 1000
	}

	//Agora estamos alterando o valor diretamente em nosso slice.
	fmt.Print(accounts) //{1100, 1200, 1300}
}
