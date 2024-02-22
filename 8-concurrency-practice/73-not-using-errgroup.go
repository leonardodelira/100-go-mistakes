package concurrencypractice

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

/*
Mistake 73: Not using errgroup

errgroup é muito útil para quando necesitamos manusear multiplas goroutines.
Por exemplo, digamos que iniciamos 3 goroutines e precisamos tratar possíveis erros que alguma dela irá retornar, como pode fazer??
Exemplo abaixo fazendo uso do errgroup
*/

/*
Fazendo uso de errgroup, não precisamos do sync.WaitGroup.
Basta criarmos nossa variavel com o contexto da goroutine pai e chamar nossas goroutines com g.Go()
Na linha 41 podemos notar que temos o g.Wait(), é aqui que temos um block até que todas as goroutines finalizem seu processo, porém
caso alguma delas tenha um erro nós conseguimos manusear o possível error também.
*/
func Mistake73_UseErrGroup(ctx context.Context, points []int) {
	results := make([]int, len(points))
	g, ctx := errgroup.WithContext(ctx)

	for i, point := range points {
		i := i
		point := point
		g.Go(func() error {
			result, err := funcwhatdoanything(ctx, point)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("success")
	fmt.Printf("results: %v", results)
}

func funcwhatdoanything(ctx context.Context, point int) (int, error) {
	if point == 3 {
		return -1, errors.New("error func: funcwhatdoanything")
	}
	return point * 2, nil
}
