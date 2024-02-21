package main

import (
	"sync"

	concurrencypractice "github.com/leonardodelira/100-go-mistakes/8-concurrency-practice"
)

func main() {
	account := concurrencypractice.Mistake70Account{
		Balances: map[string]float64{
			"fulano":   1000,
			"ciclano":  1000,
			"beltrano": 1000,
		},
	}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		account.Mistake70_AddBalance("fulano", 2000)
	}()

	go func() {
		defer wg.Done()
		account.Mistake70_BadUseMapMutex_AverageBalance()
	}()

	wg.Wait()
}
