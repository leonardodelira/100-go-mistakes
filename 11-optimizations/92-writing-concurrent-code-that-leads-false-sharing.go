package optimizations

import "sync"

/*
Mistake 92: Writing concurrent code that leads to false sharing

O exemplo fornecido utiliza duas structs, Input e Result, com a função 'count' recebendo um slice de Input e computando a soma dos campos 'a' e 'b' em Result.sumA e Result.sumB, respectivamente. Por razões ilustrativas, uma solução concorrente é implementada com uma goroutine para calcular sumA e outra para calcular sumB.

Embora esta abordagem seja válida do ponto de vista da concorrência, ela ilustra o conceito de compartilhamento falso que degrada o desempenho esperado.

Isto acontece porque as variáveis sumA e sumB são alocadas contiguamente na memória principal devido à sua posição na struct e, portanto, uma mudança em uma delas invalida a linha de cache inteira, mesmo que as atualizações sejam logicamente independentes.

Uma solução para o compartilhamento falso seria a adição de preenchimento (padding) entre as duas variáveis para garantir que elas não façam parte da mesma linha de cache, aumentando o desempenho.

Outra solução seria reestruturar o algoritmo de maneira a evitar o uso de variáveis compartilhadas, por exemplo, fazendo com que as goroutines comuniquem seus resultados locais através de canais.
*/

type Input struct {
	a int64
	b int64
}

type Result struct {
	sumA int64
	sumB int64
}

func count(inputs []Input) Result {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := Result{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
		wg.Done()
	}()

	wg.Wait()
	return result
}
