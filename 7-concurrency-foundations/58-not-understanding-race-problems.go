package concurrencyfoundations

/*
Mistake 58: Not understanding race problems

Data races vs. race conditions

Data race = quando duas ou mais goroutines acessam o mesmo local de memória.
Por exemplo:

i := 0

go func() {
    i++
}()

go func() {
    i++
}()

Nesse trecho de código acima nós temos "data race". O resultado de variavel "i" é imprevisível visto que executamos
as gorountines de forma pararela.

Umas das maneiras de resolver essa situação é por fazer uso do package sync/atomic.
var i int64

go func() {
    atomic.AddInt64(&i, 1)
}()

go func() {
    atomic.AddInt64(&i, 1)
}()

Agora garantimos que o valor de "i" sempre será dois, porque uma operação atomica não pode ser interrompida e não permite
não acessos ao mesmo tempo ao recurso.

Outra maneira de resolver a situação é por fazer uso de mutex.

i := 0
mutex := sync.Mutex{}

go func() {
    mutex.Lock()
    i++
    mutex.Unlock()
}()

go func() {
    mutex.Lock()
    i++
    mutex.Unlock()
}()
*/

/*
Race condition = quando não podemos controlar a execução dos eventos. A ordem que eles irão ser executados.
Por exemplo:
i := 0
mutex := sync.Mutex{}

go func() {
    mutex.Lock()
    defer mutex.Unlock()
    i = 1
}()

go func() {
    mutex.Lock()
    defer mutex.Unlock()
    i = 2
}()
No exemplo acima nós não temos o problema de "data race", porque estamos fazendo o uso correto de mutex.
Porém nós temos problema com "race condition", porque não sabemos em qual ordem as goroutines serão executadas.
Dessa forma é valor de "i" continua imprevisível, visto que estamos fazendo uma atribuição e não uma soma como no exemplo de data race.

*/
