package testing

import "testing"

/*
Mistake 89: Writing inaccurate benchmarks

Podemos fazer teste de BenchMark apenas por criar a funcao de teste com o prefix Bench
*/

func foo() {
	//something
}

/*
Para executar esse teste de benchmark fazemos:
go test -bench=.
*/
func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
}

/*
Caso a gente tenha que executar alguma outra função antes do benchmark, podemos resetar o timer antes de fato testar a função que queremos
*/
func BenchmarkFoo2(b *testing.B) {
	//some process or other func
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		foo()
	}
}

/*
Tambem podemos pausar o timer e despausar, pode ser util para determinado cenarios
*/
func BenchmarkFoo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		//some process or other func
		b.StartTimer()
		foo()
	}
}
