package optimizations

/*
Mistake 91: Not understanding CPU caches

"You don’t have to be an engineer to be a racing driver, but you do have to have mechanical sympathy."

CPU Cache > main memory

Memoria:
L1 = mais proximo do centro logico, 1ns para acessar o dado
L2 = 4 vezes mais lento que L1
L3 = 10 vezes mais lento que L1
*/

/*
A comparação de desempenho entre duas funções em Go é discutida neste texto. A primeira função soma todos os campos 'a' em uma fatia de structs:
*/
type Foo struct {
	a int64
	b int64
}

func sumFoo(foos []Foo) int64 {
	var total int64
	for i := 0; i < len(foos); i++ {
		total += foos[i].a
	}
	return total
}

// A segunda função também calcula uma soma, mas o argumento é uma struct contendo slices:
type Bar struct {
	a []int64
	b []int64
}

func sumBar(bar Bar) int64 {
	var total int64
	for i := 0; i < len(bar.a); i++ {
		total += bar.a[i]
	}
	return total
}

/*
Ambas as funções têm a mesma quantidade de dados, mas é discutido que a sumBar é mais rápida (cerca de 20% na máquina do autor). A principal razão para isso é uma melhor localidade espacial que faz com que a CPU busque menos linhas de cache da memória.

Usando a localidade espacial para organizar os dados de forma otimizada, podemos obter o máximo de cada linha de cache individual, melhorando a performance do programa.
Contudo, só a localidade espacial pode não ser suficiente, a previsibilidade também é um elemento crucial para otimizar o desempenho da CPU.
*/
