package datatypes

import "fmt"

/*
Mistake 20: Not understanding slice length and capacity
*/

func Test() {
	/*
		Slice faz uso de um array como estrutura de dados. Neste caso chamamos de backing array.
		O slice aponta para esse array para manusear seus valores.
		Como um array é uma estrutura de dados imutavel, nós podemos criar nosso slice e já definir um valor X de capacidade do nosso array.
		Entender a capacidade de nosso array é importante porque quando um array atinge a capacidade máxima, o slice automaticamente cria um novo array
		com o dobro de capacidade e duplica todos os elementos do array anterior.
			Isso tem um custo de processamento, e também outros fatores que serão explicados abaixo.
	*/
	//Criando um slice com length de 3 e capacidade de 6
	s := make([]int, 3, 6)

	/*
		Em memória, esse slice seria representando da seguinte maneira
			ponteiro -> [0 0 0 . . .]
		A funcao 'make' criou um slice com capacidade de 6 porém com length de 3.
			Nós podemos ter acesso apenas aos elementos que compoe o 'length' de nosso slice.
		Então como podemos fazer uso de todo capacity de nosso slice?
			Fazendo uso da função .append() do golang.
	*/
	s = append(s, 1)
	/*
		Agora nosso slice tem length de 4 e apenas dois 'slots' livres.
			ponteiro -> [0 0 0 1 . .]
		O que acontece se atingirmos o máximo do capacity do backing array?
	*/
	s = append(s, 3, 4, 5)
	/*
		Novo valor do slice: ponteiro -> [0 0 0 1 3 4]
			Não temos espaço para o número 5, nesse caso o golang irá gerar um novo backing array com mais capacidade e duplicar o array.
		Novo backing array: ponteiro -> [0 0 0 1 3 4 5 . . . . . . . . . . . .]
		Agora nosso slice aponta para esse novo backing array e o array antigo pode ser 'deletado' pelo garbage collector.
	*/

	/*
		Outro ponto de importancia é entender qual é o comportamento dos slices quando trabalhamos com slicing
	*/
	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	/*
		Neste cenário acima, ambos slices irão fazer referencia para o mesmo backing array.
		Porém o ponteiro do s2 irá começar pelo indice 1 do s1. porque é onde foi definido o slicing.
		Se uma alteração for feita em s1[0] ou s2[1], a alteração será feita no mesmo backing array e ambos os slices visualizaram o valor atualizado.

		Cuidado: os slices podem "sobreescrever" os valores um do outro por estarem apontando para o mesmo backing array.

		Se adicionarmos valores em s2 até que exceda o backing array, um novo backing array irá ser gerado para o s2 e apartir de então os slices
		não irão mais fazer referencia para o mesmo backing array.
	*/
	fmt.Print(s1, s2)
}
