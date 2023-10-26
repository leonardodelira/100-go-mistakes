package datatypes

import (
	"fmt"
	"math"
)

/*
Mistake 18: Neglecting integer overflows
Em GO nós não temos um erro quando o valor máxido do int, int32, int64 é ultrapassado, ele simplesmente volta para o valor mínimo.
Por conta disso precisamos ficar atentos quando formos trabalhar com números grandes, porque o overflow pode acontecer sem que a gente perceba.
*/

// Exemplo de um cenário que pode ocorrer um overflow, nesse caso o valor de counter vai ser o menor valor de int32 (-2147483648)
func Example1() {
	var counter int32 = math.MaxInt32
	counter++
	fmt.Printf("counter=%d\n", counter)
}

/*
Podemos fazer uso de algumas funções complementares para evitar esse problema.
No caso da função abaixo nós validamos se vai ocorrer um overflow, e se ocorrer, nós lançamos um panic.
Podemos fazer o mesmo para todas as outras operações aritméticas.
*/
func Inc32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("int32 overflow")
	}
	return counter + 1
}
