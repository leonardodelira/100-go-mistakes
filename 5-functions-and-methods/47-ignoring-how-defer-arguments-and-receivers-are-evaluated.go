package functionsandmethods

import "fmt"

/*
Mistake 47: Ignoring how defer arguments and receivers are evaluated

Devemos ter em conta que o "defer" executa no fim do método em questao, porém se passarmos uma variavel
para a função que queremos chamar com o defer, temos que ficar atentos porque o Go "armazena" a variavel
no momento em que o defer foi declarado e não o valor dela no final da execução.

Vamos ao exemplo.
*/

func BadUsageDefer() {
	i := 0
	j := 0
	defer whatIsTheValue(i) //Quando essa função ser chamada, i terá o valor de 0 e não de 1. Isso ocorre porque o Go deixa armazenado a chamada da função com o valor da variavel atual
	defer whatIsTheValue(j)

	i++
	j++
}

/*
Umas das maneiras de solucionar isso é por criar um closure.
Agora o valor de i e j será o último atribuido a ele.

Outra maneira também é por passar como parametro para nosso metodo de destino o ponteiro da variavel.
Dessa forma não precisamos fazer uso de defer.
*/
func GoodUsageDefer() {
	i := 0
	j := 0
	defer func() {
		whatIsTheValue(i)
		whatIsTheValue(j)
	}()

	i++
	j++
}

func whatIsTheValue(n int) {
	fmt.Print(n)
}
