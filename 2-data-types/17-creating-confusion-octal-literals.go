package datatypes

import "fmt"

/*
Mistake 17: Creating confusion with octal literals
*/

/*
O que será que a função abaixo irá printar?
*/
func example() {
	sum := 100 + 010
	fmt.Println(sum)
}

/*
Resulta de example() é 108. Em Golang um integer que começa com 0 é considedo um octal integer (base 8).
E o valor de 010 em base 8 é = 8. Por isso o resultado de 108.
O importante aqui é ficar atento a quando temos um octal integer, porque nesse caso podemos deixar ele mais explicito utilziando
o caracter "o".
*/

/*
Neste exemplo abaixo deixamos explicto para um futuro leitor do código que se trata de um octal integer.
*/
func example2() {
	sum := 100 + 0o10
	fmt.Print(sum)
}
