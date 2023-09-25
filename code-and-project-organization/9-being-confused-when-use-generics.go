package codeandprojectorganization

/*
Mistake 9: Being confused about when to use generics

A chave é não usar generics prematuramente e considerá-los apenas quando você está prestes a escrever código repetitivo. As generics são uma forma de abstração, e o uso inadequado de abstrações pode introduzir complexidade desnecessária no código.

Casos Comuns para o Uso de Generics em Go (Quando Usar):

- Em estruturas de dados, como árvores binárias, listas encadeadas, ou heaps.
- Em funções que trabalham com slices, maps e canais de qualquer tipo.
- Para fatorar comportamentos (por exemplo, na biblioteca sort).

Quando Não Usar Generics em Go (Quando Não Usar):

- Quando você está chamando um método do tipo argumento diretamente.
- Quando a introdução de generics torna o código mais complexo e menos legível.



Exemplos:
Trabalhando com Slices Genéricos:
package main

import (
	"fmt"
)

// Função para encontrar o maior elemento em um slice genérico
func FindMax[T comparable](items []T) T {
	max := items[0]
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	return max
}

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	maxInt := FindMax(intSlice)
	fmt.Println("Maior inteiro:", maxInt) // Saída: Maior inteiro: 9

	floatSlice := []float64{3.14, 1.0, 2.71, 0.5}
	maxFloat := FindMax(floatSlice)
	fmt.Println("Maior float:", maxFloat) // Saída: Maior float: 3.14
}


Fatorando Comportamentos com Generics:
package main

import (
	"fmt"
	"sort"
)

// Função genérica para ordenar qualquer tipo de slice
func SortSlice[T comparable](items []T) {
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
}

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	SortSlice(intSlice)
	fmt.Println("Slice ordenado:", intSlice) // Saída: Slice ordenado: [1 1 2 3 3 4 5 5 5 6 9]

	strSlice := []string{"banana", "apple", "cherry", "date"}
	SortSlice(strSlice)
	fmt.Println("Slice ordenado:", strSlice) // Saída: Slice ordenado: [apple banana cherry date]
}


*/
