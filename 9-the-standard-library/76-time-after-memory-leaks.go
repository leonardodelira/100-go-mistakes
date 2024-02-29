package thestandardlibrary

import (
	"fmt"
	"time"
)

/*
Mistake 76: time.After and memory leaks

Quando chamamos time.After devemos lembrar que o recurso alocado para essa chamada só vai ser liberada quando
o timer expirar. Ou seja se criamos um timer.After(1 hora) o recurso alocado só será liberado depois de 1 hora.
Por conta disso nós podemos ter "memomy leaks" se usarmos esse recurso de maneira ineficiente.
Se fazemos uso de timer.After em chamadas https, loops, consumers de kafka (ou qualquer outro) podemos estar deixando nossa
aplicação ineficiente, isso porque novas instancias de timer.After será criada a cada iteração.

Se precisamos por alguma razão contar o tempo para realizar alguma ação, podemos fazer uso de time.NewTimer()
*/

func Mistake76_consumer(ch chan int) {
	timerDurantion := time.Second * 5
	timer := time.NewTimer(timerDurantion)

	for {
		select {
		case event := <-ch:
			fmt.Printf("chegou o número %d ", event)
			timer.Reset(timerDurantion)
		case <-timer.C:
			fmt.Print("já se passaram 5 segundos sem nenhum dado novo")
		}
	}
}
