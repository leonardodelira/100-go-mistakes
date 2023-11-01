package datatypes

import "fmt"

/*
Mistake 21: Inefficient slice initialization
*/

type Foo struct {
	a string
	b string
}

type Bar struct {
	c string
	d string
}

func fooToBar(foo Foo) Bar {
	return Bar{
		c: foo.a,
		d: foo.b,
	}
}

/*
Dependendo do cenario, não passar corretamento o length e/ou o capacity na criação do nosso slice pode ser uma má pratica.
Vejamos o exemplo abaixo.
Digamos que tenhamos uma funcao que faça uma conversão de um struct Foo para uma struct Bar
*/

/*
Neste exemplo, nós criamos um slice de bars sem difinir um tamanho de length e nem de capacity.
O que vai ocorre aqui é que todas as vezes que o backing array de bars chegar ao limite, o golang
terá que criar um novo backing array e copiar os elementos para esse novo array.
Dependendo do tamanho de Foo, esse cópia e geração de um novo backing array pode ocorrer várias vezes. E isso
prejudica no tempo de execução do código.
*/
func Example1_Mistake21(foos []Foo) {
	bars := make([]Bar, 0)

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}

	fmt.Print(bars)
}

/*
Temos duas soluções para esse cénario, onde podemos definir o tamanho de length e do capacity
*/

/*
Neste exemplo abaixo definimos o tamanho do capacty. Dessa forma o golang já vai alocar um backing array suficiente para a operação
(Considerando é claro que o tamnho de foos e bars serão o mesmo.)
Dessa forma novos backing array não serão gerados e evitara perda de processamento e custo para o GC.
*/
func Example2_Mistake21(foos []Foo) {
	bars := make([]Bar, 0, len(foos))

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}

	fmt.Print(bars)
}

/*
Esta outra abordagem também é eficiente, nós já criamos o nosso slice com length suficiente.
O que muda é a forma que atualizamos os valores, acessando diretamente o index do slice já que criamos o slice com todos os default preenchidos.
Dessa forma novos backing arrays não serão criados.

Obs:
O example3 executa o código mais rápido que example2 por alterar diretamente o index do slice.
O que irá definir qual abordagem vamos preferir será o contexto do nosso código, muitas vezes trabalhar com o append facilita
a leitura do código.
*/
func Example3_Mistake21(foos []Foo) {
	bars := make([]Bar, len(foos))

	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}

	fmt.Print(bars)
}
