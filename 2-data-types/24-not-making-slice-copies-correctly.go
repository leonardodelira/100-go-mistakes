package datatypes

import "fmt"

/*
Mistake 24: Not making slice copies correctly

O golang nos forcene uma função chama copy() para realizer cópias de um slice a outro.
Porém temos que ficar atentos na maneira que fazemos uso.
A função copy() irá copiar o "minimo de elmentos com base no lenght do slice de destino".
*/

func ExampleMistake24() {
	origin := []int{1, 2, 3, 4, 5}
	var dest []int
	copy(dest, origin)

	//o slice 'dest' neste caso terá valor de []
	//isso ocorre porque a função copy irá copiar o minimo de elementos entre origin e dest
	//neste caso o minimo é o 'dest' que tem lenght de 0, visto que o lenght de origin é 5
	fmt.Print(dest)
}

func ExampleMistake24_Correct() {
	origin := []int{1, 2, 3, 4, 5}
	var dest = make([]int, len(origin))
	copy(dest, origin)

	//Agora a cópia será feita de modo correto. O lenght do destino é igual (ou poderia ser maior) que a origem
	fmt.Print(dest)
}
