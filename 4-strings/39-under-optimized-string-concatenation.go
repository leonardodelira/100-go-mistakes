package strings

import "strings"

/*
Mistake 39: Under-optimized string concatenation

Devemos ficar atentos ao fazer concatenação de strings em go.
Em go uma string é imutável, então dependendo da forma que fazemos a concatenação podemos gerar várias strings na memória.
e por conta disso perder performance.
*/

func BadExampleConcatString(values []string) {
	s := ""
	for _, value := range values { //Estamos alocando em memória varias variaveis porque string é imutavel.
		s += value
	}
}

/*
Essa é uma maneira mais eficiente de concatenar strings. O método `WriteString` da struct `Builder`
é usado para anexar strings ao buffer interno, evitando cópias desnecessárias de memória.
Porém ainda temos um problema aqui, nosso buffer interno tem um limite pre alocado de valores que pode receber
se esse valor for ultrapassado uma nova relocação e copia de valores terá que ser feita e isso é custoso.
- Aqui é aquele ponto que já falamos sobre backing array e slices.
Então a solução para tal problema é definir quantos de espaço em memória nosso struct builder vai precisar com base no input
*/
func GoodExampleConcatString(values []string) {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
}

func BetterGoodExampleConcatString(values []string) {
	total := 0
	for i := 0; i < len(values); i++ { //Aqui estamos somando o total de todos os bytes de cada string
		total += len(values[i])
	}

	sb := strings.Builder{}
	sb.Grow(total) //Agora dizemos ao Go para pre alocar em memoria a quantia exata de bytes, dessa forma evitamos a criaçao de um novo backing array e a cópia de todos os valores.
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
}
