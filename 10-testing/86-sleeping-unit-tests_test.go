package testing

import (
	"testing"
)

/*
Mistake 86: Sleeping in unit tests
*/

type MockPublisher struct {
	ch chan []Foo
}

func (p *MockPublisher) Publish(foos []Foo) {
	p.ch <- foos
}

func TestGetBestFoo(t *testing.T) {
	foos := []Foo{
		{
			Act: 1,
		},
		{
			Act: 2,
		},
		{
			Act: 3,
		},
	}

	mockPublisher := &MockPublisher{
		ch: make(chan []Foo),
	}
	defer close(mockPublisher.ch)

	handler := &Handler{
		N:         2,
		Publisher: mockPublisher,
	}

	//Para testar o que a função retorna é simples, mas e para testar o que foi executado na goroutine?
	got := handler.getBestFoo(foos)
	expected := foos[0].Act
	if got.Act != expected {
		t.Errorf("expected: %d, got: %d", expected, got.Act)
	}

	/*
		Aqui vemos como testar o que foi passado na goroutine, uma das maneiras é por trabalhar com sincronização fazendo
		uso de um channel.
		Dessa maneira podemos validar que a função que foi chamada dentro da gorutine recebeu os valores corretos.
	*/
	if v := len(<-mockPublisher.ch); v != handler.N {
		t.Errorf("expected: %d foos, got: %d", len(foos), len(<-mockPublisher.ch))
	}
}
