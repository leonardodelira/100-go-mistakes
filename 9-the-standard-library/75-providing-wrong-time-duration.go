package thestandardlibrary

import (
	"time"
)

/*
Mistake 75: Providing a wrong time duration

Em Go, para evitar qualquer tipo de erro ou confusão ao trabalhar com "time", podemos fazer uso da API time.Duration.
*/

// Fazenso uso da api, estamos seguros que nosso ticket vai ser de 1 segundo.
func Mistake75() {
	ticker := time.Duration(time.Second)
	time.Sleep(ticker)
}
