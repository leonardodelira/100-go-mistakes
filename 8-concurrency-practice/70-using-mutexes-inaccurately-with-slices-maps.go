package concurrencypractice

import (
	"fmt"
	"sync"
)

/*
Mistake 70: Using mutexes inaccurately with slices and maps

Assim como os slices, os maps também fazendo referencia para o que chamamos de dataset.
É nesse dataset que temos alguns metadados do nosso map, por exemplo um counter e os ponteiros para os buckets.
Dito isso, devemos ficar atentos quando vamos trabalhar com mutação + concorrencia envolvendo maps.
Exemplo abaixo.
*/

type Mistake70Account struct {
	mu       sync.RWMutex
	Balances map[string]float64
}

/*
Se duas goroutines estiveram sendo executadas e ambas tentar acessar nossa struct que tem o map balances poderemos ter um
problema de data race.
Isso ocorre porque como dito, o map possui uma especie de "backing array", assim como nas structs.
Então mesmo que a gente copie os dados do map para uma nova variavel (como é feito na linha 31)
podemos sofrer com data race, porque duas goroutines estão tentando alterar o dataset do map.
*/
func (a *Mistake70Account) Mistake70_BadUseMapMutex_AverageBalance() {
	a.mu.Lock()
	balances := a.Balances
	a.mu.Unlock()

	sum := 0.
	for _, v := range balances {
		sum += v
	}
	fmt.Println(sum)
}

func (a *Mistake70Account) Mistake70_AddBalance(id string, balance float64) {
	a.mu.Lock()
	a.Balances[id] += balance
	a.mu.Unlock()
}

/*
Para solucionar isso, temos duas abordagens:
1 - Se o processamento da nossa função for rapida, como por exemplo na função da Average (linha 29). Podemos manter o lock durante
toda a execução da função. (Não apenas para criar uma cópia do map).

2 - A segunda opção que temos é criar uma cópia do map, porém da maneira correta. Para que a nova variavel não aponte para o mesmo
dataset do map original. Exemplo abaixo:
*/

// Agora temos a variavel balances sem apontar para o mesmo dataset do map original.
func (a *Mistake70Account) Mistake70_GoodUseMapMutex_AverageBalance() {
	a.mu.Lock()
	balances := make(map[string]float64, len(a.Balances))
	for k, v := range a.Balances {
		balances[k] = v
	}
	a.mu.Unlock()

	sum := 0.
	for _, v := range balances {
		sum += v
	}
	fmt.Println(sum)
}
