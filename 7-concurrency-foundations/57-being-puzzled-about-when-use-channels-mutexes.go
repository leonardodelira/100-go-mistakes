package concurrencyfoundations

/*
Mistake 57: Being puzzled about when to use channels or mutexes

channels = normalmente usamos para concurrency goroutines.
mutexes = normalmente usamos para parallel goroutines.

channels = passamos valores entre goroutines.
mutexes = garantimos que um determinado recurso possa ser usado por goroutines pararelas. Em geral gorountines pararelas precisam
estar sincronizadas, por isso se elas compartilham algum recurso, por exemplo uma struct. Fazemos o uso de mutexes para essa
coordenação.
*/
