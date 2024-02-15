package concurrencypractice

import (
	"fmt"
	"sync"
)

/*
Mistake 63: Not being careful with goroutines and loop variables

Devemos ficar atentos ao usar goroutines dentro de loops, e passar a variavel do loop para a goroutine.
*/

/*
Neste exemplo, o resulto do fmt.print() é indeterministico. Quando a goroutine executar o valor de i pode ser outro, não mais
o mesmo de quando nós criamos a goroutine.
Ou seja, quando executarmos a goroutine com i=0, talvez quando a goroutine ser executada o valor de i já seja 2
*/
func ExampleErrorGoroutineLoop() {
	sg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		sg.Add(1)
		go func() {
			defer sg.Done()
			fmt.Print(i)
		}()
	}
	sg.Wait()
}

/*
Podemos solucionar essa questao de duas formas, criando uma variavel local dentro do for. Ou passando
o valor de i como parametro para nosso closure function que executa a goroutine.
No nosso caso, vamos fazer a segunda opção.

Agora 0,1,2 é printado na tela .. é claro que a ordem do print é inderteministica porque a execução das goroutines varia.
Mas podemos estar seguros que todos os números serão mostrados na tela.
*/
func ExampleCorrectGoroutineLoop() {
	sg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		sg.Add(1)
		go func(val int) {
			defer sg.Done()
			fmt.Print(val)
		}(i)
	}
	sg.Wait()
}
