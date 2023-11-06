package datatypes

import (
	"encoding/json"
	"fmt"
)

/*
Mistake 22: Being confused about nil vs. empty slices
*/

/*
A slice is empty if its length is equal to 0.
A slice is nil if it equals nil.
*/

func ExampleNilEmptySlices() {
	var a []string
	b := []string{}
	c := make([]string, 0)

	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
}

/*
O slice 'a' será: empty, nil
O slice 'b' será: empty, non-nil
O slice 'c' será: empty, non-nil

Importante lembrar que a opçao do slice 'a' não irá criar alocação em memória, porque está como nil.

Quando usar cada uma das opções?
Opção slice 'a':
	Quando não sabemos o tamanho que terá o slice, ou se quer que terá valor. Exemplo abaixo.
	No exemplo abaixo nós podemos ter alguma valor no slice 'a', porém também pode ocorrer de nunca ter valor algum. Nesse caso
	criar o slice como empty, nil é vantajoso para não alocar nada em memória, já que não será usado.
*/

func foo() bool { return true }
func bar() bool { return true }

func ExampleSliceA() []string {
	var a []string

	if foo() {
		a = append(a, "foo")
	}

	if bar() {
		a = append(a, "bar")
	}

	return a
}

/*
Opção slice 'b':
Devemos declarar nossos slices dessa maneira apenas quando tivermos a intenção de criar o slice com valores iniciais.
*/
func ExampleSliceB() {
	b := []string{"foo", "bar"}
	fmt.Print(b)
}

/*
Opção slice 'c':
Podemos declarar dessa maneira quando sabemos o tamanho final que terá nosso slice.
Dessa forma evitamos copias do backing array.
*/
func ExampleSliceC(ints []int) {
	c := make([]string, len(ints))
	fmt.Print(c)
}

/*
Um ponto a se considerar: algumas lib do golang diferenciam nil de empty.
Como exemplo o enconding/json.
Ao fazer o Marshal de um valor nil, a lib irá retornal null
E ao fazer Marshal de um valor empty, a lib irá retornar []
*/

type customer struct {
	ID         string
	Operations []float32
}

func ExempleFinalSliceNilDistinct() {
	var s1 []float32
	customer1 := customer{
		ID:         "foo",
		Operations: s1,
	}
	b, _ := json.Marshal(customer1)
	fmt.Println(string(b)) //Após a operação de marshal, o json de "operations" será null

	s2 := make([]float32, 0)
	customer2 := customer{
		ID:         "bar",
		Operations: s2,
	}
	b, _ = json.Marshal(customer2)
	fmt.Println(string(b)) //Após a operação de marshal, o json de "operations" será []
}
