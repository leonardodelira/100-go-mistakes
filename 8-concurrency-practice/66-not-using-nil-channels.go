package concurrencypractice

import (
	"fmt"
	"sync"
	"time"
)

/*
Mistake 66: Not using nil channels

Deixar o channel como "nil" pode ser útil para sabermos quando parar de "escutar" esse channel.
Por exemplo, um "chan int" quando é fechado, ou seja, quando chamados a função close(chan) irá sempre enviar para o reciver o valor de 0.
Isso ocorre porque 0 é o valor padrão de int.
Nesse caso podemos ter um vazamento de memória ou até mesmo um bug por estar escutando um channel que já foi fechado e fica enviando sempre
o mesmo valor
Exemplo abaixo.
*/

// Exemplo de um channel que é fechado (close) e mesmo assim continuamos a escutar ele na função is_running()
func TestExample1Mistake66() {
	c := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go closeChannelAfterLoop(c)
	go is_running(c)
	wg.Wait()
}

// Aqui fechamos o channel
func closeChannelAfterLoop(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}

/*
Aqui estamos escutando o channel, todos os valores novos que chega.
Acontece que em determinado momento, na função de cima o channel será fechado mas mesmo assim o nosso for/select
continuara a escutar o channel e ficará recebendo o valor de 0.
*/
func is_running(ch1 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		}
	}
}

/*
Umas das possíveis soluções para isso é por fazer uso do boolean que nos diz que o channel continua aberto.
Além disso, depois que recebermos que o channel está fechado, podemos definilo como "nil", dessa forma o select statment não
irá mais ficar esperando por nenhuma novidade do channel.
E quando os dois channel forem nil, o for loop não será mais executado.
*/
func is_running2(ch1, ch2 chan int) {
	for ch1 != nil || ch2 != nil {
		select {
		case v, open := <-ch1:
			if !open {
				ch1 = nil
				break
			}
			fmt.Println(v)
		case v, open := <-ch2:
			if !open {
				ch2 = nil
				break
			}
			fmt.Println(v)
		}
	}
}

/*
Lembrando que essa abordagem é útil para quando não queremos usar o range para iterar sobre o channel.
	Quando fazemos uso do range para iterar em um channel e o channel é fechado o range quebra (break).
Em determinadas abordagens queremos ter um controle melhor sobre quando o channel foi fechado.
*/
