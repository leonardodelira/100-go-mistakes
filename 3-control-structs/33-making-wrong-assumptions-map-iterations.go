package controlstructs

import "fmt"

/*
Mistake 33: Making wrong assumptions during map iterations

- Ordering
Importante lembrar que um map não garante a ordem os elementos, ou seja, se inserimos o valor A e depois o B não podemos ter a certeza
que eles ficam em ordem dentro do map.
	Com isso em mente não podemos presumir que iterar sobre eles, teremos sempre a mesma ordem.

- Map update during an iteration
Atualizar um map enquanto iteramos sobre ele também pode resultar em resultados nao deterministicos, ou seja, variados.
*/

func ExampleMapUpdateIteration() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	/*
		Qual será o resultado desse map? A cada iteração irá variar, exemplos abaixo:
		map[0:true 1:false 2:true 10:true 12:true 20:true 22:true 30:true]
		map[0:true 1:false 2:true 10:true 12:true 20:true 22:true 30:true 32:true]
		map[0:true 1:false 2:true 10:true 12:true 20:true]
		Isso ocorre porque a linguaguem do Go funciona assim, como desenvolvedores não podemos forçar para ser diferente.
			O fato é que ler o map e altera-lo ao mesmo tempo (inserir ou deletar) vai gerar esses resultados imprevisiveis.
			A solução está na outra função.
	*/
	fmt.Println(m)
}

// Para soluciionar o caso acima, podemos criar uma cópia do map original e alterar a cópia.
// Dessa forma estaremos lendo o map original mas fazendo updates na cópia.
func ExampleMapUpdateIterationCorrect() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	m2 := copyMap(m)

	for k, v := range m {
		m2[k] = v
		if v {
			m2[10+k] = true
		}
	}

	fmt.Println(m2)
}

func copyMap(m map[int]bool) map[int]bool {
	newMap := make(map[int]bool, len(m))
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
