package concurrencypractice

import (
	"fmt"
	"sync"
	"time"
)

/*
Mistake 72: Forgetting about sync.Cond

sync.Cond é uma maneira de notificar varias goroutines a que determinado evento ocorreu.
Dessa forma não ficamos escutando o tempo todo, como por exemplo fariamos com um channel.
*/

type Donation struct {
	balance int
	cond    sync.Cond
}

func Mistake72_SyncCond() {
	donation := &Donation{
		balance: 0,
		cond:    *sync.NewCond(&sync.Mutex{}),
	}

	go donation.Goal(10)
	go donation.Goal(15)

	for {
		time.Sleep(1 * time.Second)
		donation.cond.L.Lock()
		donation.balance++
		fmt.Println(donation.balance)
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}

/*
Fazendo uso do sync.Cond, podemos "dispensar" o processamento da goroutine até que um Broadcasting chegue.
Na nossa função principal depois de cada update nós enviamos um Broadcasting, dessa forma todas as goroutines
que estão acessando a mesma struct escutam esse evento e continuam a processar a partir do último ponto que parou
que seria logo após o "Wait()".
*/
func (d *Donation) Goal(goal int) {
	d.cond.L.Lock()
	for d.balance < goal {
		d.cond.Wait()
	}
	fmt.Printf("Atingimos a meta de: %d\n", goal)
	d.cond.L.Unlock()
}
