package datatypes

import "fmt"

/*
Mistake 27: Inefficient map initialization

Um Map em Golang tem um sistema de hash table por trás.
Quando nós iniciamos nosso Map sem dizer seu "size" e recebemos muitos elementos, o map terá que crescer e fazer um rebalanceamento
de seus indexs.
	Esse processo é custoso, por isso aconselha-se iniciar o Map com o size pre definido, se for possível.
*/

func InitMapWithSize() {
	m := make(map[string]int, 1000000)
	fmt.Print(m)
}

/*

Exemplo que o livro trás de um benchmark
BenchmarkMapWithoutSize-4     6    227413490 ns/op
BenchmarkMapWithSize-4       13     91174193 ns/op

Como podemos ver, a segunda opção é 60% mais rapida, porque a inicialização do Map foi feita pre definindo seu size.
*/
