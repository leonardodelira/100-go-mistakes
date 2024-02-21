package concurrencypractice

import (
	"fmt"
	"sync"
)

/*
Mistake 69: Creating data races with append

Quando criamos um slice, devemos lembrar que temos um backing array por de trás desse slice.
Ao realizar a operação de append, Go verifica se o backing array ainda tem capacity para adicionar o item.
Caso tenha espaço, apenas incrementa um elemento novo no backing array, caso não tenha espaço gera um novo backing array com mais
capacity para aquele slice.

Não teremos problema de data race se nosso slice estiver cheio e duas goroutines tentar incrementar um novo elemento (desde que seja em um novo slice).
Porque o Go irá alocar um novo backing array para esse slice novo.

Porém teremos problemas se duas ou mais goroutines realizar o append em um slice que não está cheio, porque ambas as goroutines
irão tentar atualizar o mesmo indice do backing array, nesse caso teremos data race.
*/

/*
Aqui teremos um problema de data race no append, mesmo que estejamos criando um novo slice.
Como já vimos, o backing array do slice é mantido pelo Go .. a não ser que o lenght seja igual ao capacity.
No caso abaixo, criamos o slice "s" com lenght = 0 e capacity = 1, nesse caso o Go irá tentar incrementar
um novo valor no mesmo backing array (mesmo que estejamos criando um novo slice).
Por conta disso, as duas goroutines irá tentar incrementar um valor no mesmo indice do slice, dessa forma estaremos sofrendo de
data race.

Uma possível solução é criar um cópia do slice original para não ficarmos vinculados ao mesmo backing array.
sCopy := make([]int, len(s), cap(s))
copy(sCopy, s)
s1 := append(sCopy, 1)
*/
func Mistake69_AppendDataRace() {
	s := make([]int, 0, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	go func() {
		s1 := append(s, 1)
		fmt.Println(s1)
		wg.Done()
	}()

	go func() {
		s2 := append(s, 2)
		fmt.Println(s2)
		wg.Done()
	}()
}

/*
Já neste examplo não sofremos com data race no append, porque visto que o slice foi criado com lenght = 1 e capacity = 1. Quando
o Go tentar realizar um novo append ele irá notar que o backing array está cheio e precisa ser criado um novo.
Dianta disso cada novo slice "s1" e "s2" terá seu próprio backint array.
*/
func Mistake69_NotOccurAppendDataRace() {
	s := make([]int, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	go func() {
		s1 := append(s, 1)
		fmt.Println(s1)
		wg.Done()
	}()

	go func() {
		s2 := append(s, 2)
		fmt.Println(s2)
		wg.Done()
	}()
}
