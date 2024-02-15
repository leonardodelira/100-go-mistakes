package concurrencypractice

import (
	"fmt"
	"sync"
)

/*
Mistake 64: Expecting deterministic behavior using select and channels

Quando fazemos o uso de "select" para escutar channels, devemos ter em mente que a ordem que colocamos os "cases" não
necessáriamente será a prioridade que o Go levará em conta para executar os comandos respectivos.
Se chegar um evento para dois channels que estão no select, o Go de forma aleatória irá executar um deles.
*/

func ProducerMessages() {
	messageCh := make(chan int, 10)
	disconectCh := make(chan interface{})

	sw := sync.WaitGroup{}
	defer sw.Wait()

	sw.Add(1)
	go consumerOneSolution(messageCh, disconectCh, &sw)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}

	disconectCh <- struct{}{}
}

/*
Neste exemplo abaixo podemos não sabemos se a execução do processo chegará ao fim. Porque se chegar eventos nos dois
channels ao mesmo tempo, Go irá executar um dos "cases" de forma aleatória.
*/
func consumerWrong(messagech chan int, disconectCh chan any, sw *sync.WaitGroup) {
	for {
		select {
		case v := <-messagech:
			fmt.Println(v)
		case <-disconectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

/*
Uma possível solução é, antes de desconectar o channel .. podemos verificar se todas as mensagens do outro channel foi lida.

Outra solução também seria criar o channel como "unbuffered". No caso nosso channel tem buffer de 10
*/
func consumerOneSolution(messagech chan int, disconectCh chan any, sw *sync.WaitGroup) {
	defer sw.Done()

	for {
		select {
		case v := <-messagech:
			fmt.Println(v)
		case <-disconectCh:
			for {
				select {
				case v := <-messagech:
					fmt.Println(v)
				default:
					fmt.Println("disconnection, return")
					return
				}
			}
		}
	}
}
