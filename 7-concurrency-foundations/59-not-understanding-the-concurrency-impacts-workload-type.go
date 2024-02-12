package concurrencyfoundations

import (
	"fmt"
	"runtime"
)

/*
Mistake 59: Not understanding the concurrency impacts of a workload type

Temos dois tipos de "workload" principais.
CPU = processamento lógico da nossa aplicação, por exemplo um algoritmo de merge sort
I/O = acesso ao bancos de dados ou uma chamada REST.

Ao trabalhar em uma aplicação GO com concorrencia onde o workload é I/O, o ideal é saber se a aplicação que vamos
chamar suporta chamadas simultaneas.

Porém, se nosso workload é com CPU devemos criar nossas goroutines de forma consciente.
Por exemplo, podemos criar um "pool" de goroutines baseado no quantidades de CPUs fisicos disponiveis na maquina que estamos
executando o código. Para isso podemos fazer uso de GOMAXPROCS.

Devemos lembrar que se criarmos goroutines demais, iremos aumentar o context switch entre elas. E isso pode ocasionar
em perda de desempenho da nossa aplicação.

Por isso que: se criarmos as goroutines com base na quantidade de CPUs disponiveis na maquina, faremos melhor uso dos recursos.
*/

func Gomaxprocs() {
	n := runtime.GOMAXPROCS(0)
	fmt.Print(n)
}
