package concurrencypractice

/*
Mistake 71: Misusing sync.WaitGroup

Ao fazermos uso de waitGroup, devemos ficar atentos a como utilizamos a funcao wg.Add()
Devemos executar essa função sempre na goroutine pai. Ou seja, normalmente é a main.

Por exemplo, o trecho de código abaixo está errado. Não é correto chamar o wg.Add() dentro da goroutine que estamos criando
wg := sync.WaitGroup{}
var v uint64

for i := 0; i < 3; i++ {
    go func() {
        wg.Add(1)
        atomic.AddUint64(&v, 1)
        wg.Done()
    }()
}

wg.Wait()
fmt.Println(v)
*/
