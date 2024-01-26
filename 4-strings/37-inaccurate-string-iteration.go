package strings

import (
	"fmt"
	"unicode/utf8"
)

/*
Mistake 37: Inaccurate string iteration

Devemos lembrar que em uma string com caracteres especiais, esse caracter especial pode ser composto por mais de um byte.
Quando fazemos uma iteração sobre essa string com caracter especial e acessamos seu index, na verdade estamos acessando
o primeiro byte do conjunto que forma o caracter especial, e isso pode ocasionar um erro.
*/

/*
O resultado desse range e len() será:
position 0: h
position 1: Ã
position 3: l
position 4: l
position 5: o
len=6
Porque? Porque o caracter "ê" é composto por mais de um byte, e quando acessamos o index da string "s[1]" estamos acessando o primeiro byte
do conjunto de bytes. E esse byte sozinho tem o valor de "ã".
Então por exemplo:
"ê" = bytes[x001, x002]
"ã" = bytes[x001]
Como dito, se acessamos direto o index vamos acessar o primeiro byte do conjuto de bytes .. por isso outra letra foi mostrada.
E a razão pela qual o len = 6, é porque a função len() conta o números de bytes da string .. nesse caso um dos caracteres tem mais de um byte.
*/
func IterateStringWrong() {
	s := "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))
}

/*
Forma correta:
Utilizar a variavel que o range cria ao iterar sobre a string vai nos mostrar o valor correto da rune
E utilizar RuneCountInString vai contar os caracteres da string e não os bytes.
*/
func IterateStringCorrect() {
	s := "hêllo"
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}
	fmt.Printf("len=%d\n", utf8.RuneCountInString(s))
}

/*
Outra forma de acessar um index especifico de nossa string é por criar um slice de runes antes.
Devemos usar essa opção com cuidado considerando que criamos um novo slice na memória e fazemos a conversão da string.
*/
func IterateStringCorrect2() {
	s := "hêllo"
	runes := []rune(s)
	for i := range runes {
		fmt.Printf("position %d: %c\n", i, runes[i])
	}
	fmt.Printf("len=%d\n", utf8.RuneCountInString(s))
}
