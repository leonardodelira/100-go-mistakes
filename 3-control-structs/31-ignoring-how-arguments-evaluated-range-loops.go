package controlstructs

import "fmt"

/*
Mistake 31: Ignoring how arguments are evaluated in range loops

Devemos ter atenção ao usar range ou um for loop convencional.
Quando fazemos uso de "range" o valor é atribuido apenas uma vez.
Porém quando fazemos uso de for, a cada intereção o valor de lenght é alterado.
Vejamos um exemplo abaixo
*/

/*
Neste exemplo abaixo nós poderiamos pensar que o "for range" nunca teria fim, visto que estamosa adicionando elementos a cada loop.
Porém o que acontece nesse caso é que quando fazemos uso de um range, o golang cria uma variavel temporaria como copia da variavel original,
dessa forma depois que ele percorrer os 3 elementos do slice o for irá terminar.
*/
func ExampleRange() {
	s := []int{0, 1, 2}
	for range s {
		s = append(s, 10)
	}
}

/*
Neste caso o nosso for nunca terá fim, porque a cada iteração do for o "len" será executado novamente, e ele continuará crescendo.
Assim o valor de "i" sempre será menor que o "len" do slice "s"
*/
func ExampleFor() {
	s := []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		s = append(s, 10)
	}
}

/*
Com array também temos um comportamento diferente: quando usamos o "range" o golang cria uma cópia do array original.
Se quisermos alterar o valor do array ou algo coisa do tipo, devemos ficar atentos se estamos acessando a cópia ou o valor original.
No exemplo abaixo estamos alterando o último valor do array para 10, mas quando printamos o valor continuará a ser 3. Porque nesse caso
alteramos o valor do array original mas estamos mostrando o valor da cópia que o "range" criou.
*/
func ExampleArray() {
	s := [3]int{1, 2, 3}
	for i, v := range s {
		s[2] = 10
		if i == 2 {
			fmt.Print(v) //3
		}
	}
}

// Se quisermos acessar o valor atualizado, devemos acessar diretamente com o index
func ExampleArray2() {
	s := [3]int{1, 2, 3}
	for i := range s {
		s[2] = 10
		if i == 2 {
			fmt.Print(s[i]) //10
		}
	}
}
