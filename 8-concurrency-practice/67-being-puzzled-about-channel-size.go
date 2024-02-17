package concurrencypractice

import "fmt"

/*
Mistake 67: Being puzzled about channel size

unbufered channel = quando queremos ter sincronização entre nossas goroutines
buffered channel = são casos raros, por exemplo quando temos que ter algum tipo de rate-limit (sinceramente não entendi muito bem a razão de usar un buffered channel).

quando fazemos uso de um unbeffered channel, nós temos a garantia de sincronização. Isso porque teremos um block até que o
"sender" tenha certeza que o "receiver" recebeu a mensagem.

diferentemente do buffered channel que não espera garantia nenhuma, se tem espaço no buffer do channel o dado é enviado e tchau.
*/

/*
No caso abaixo temos um channel unbuffered, ou seja, o sender precisa receber uma garantia de leitura para continuar o processo
O que acontece nesse caso abaixo é: estamos atribuindo um valor ao channel, porém em nenhum lado estamos fazendo a leitura desse valor
por conta disso temos um block no processo. Mesmo tendo a linha 26 que na teoria deveria printar o valor, o processo nunca chega
na linha 26 porque está block na linha 25.
*/
func Mistake67_TestUnbufferedChannel() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}

/*
Neste caso abaixo é diferente, como dito um channel buffered não espera a garantia de leitura de um valor que entrou no channel
Ele passa o valor e continua o processo.
Por isso a execução da função seria completa.
*/
func Mistake67_TestBufferedChannel() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)
}
