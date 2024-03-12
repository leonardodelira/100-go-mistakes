package testing

/*
Mistake 86: Sleeping in unit tests

Para testar funções que tem chamadas de goroutines, podemos usar algumas estratégias para ser mais exato
em nossos testes.
Podemos fazer uso de retry ou sincronização.
Abaixo temos a função e no arquivo de teste a maneira correta de testar.
*/

type Foo struct {
	Act int
}

type Publisher interface {
	Publish([]Foo)
}

type Handler struct {
	N         int
	Publisher Publisher
}

func (h *Handler) getBestFoo(someFoos []Foo) Foo {
	foos := someFoos
	best := someFoos[0]

	go func() {
		if h.N < len(foos) {
			foos = foos[:h.N]
		}
		h.Publisher.Publish(foos)
	}()

	return best
}
