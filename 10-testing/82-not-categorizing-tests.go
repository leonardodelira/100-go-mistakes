package testing

import (
	"os"
	"testing"
)

/*
Mistake 82: Not categorizing tests

Em go podemos criar tags para categorizar o tipo de test que temos. Com as tags nós podemos executar
apenas determinado tipo de teste, com isso ganhamos tempo e qualidade em nossos tests.
Para criar uma tag, basta escrever o seguinte no topo do arquivo de test:

"//go:build integration"

Nesse caso a tag se chama "integration", mas poderia ser qualquer outra.

Depois para executar os testes que tem tag basta indicar para o Go qual tag queremos executar.

## go test --tags=integration -v .
*/

/*
Porém podemos ter uma desvantagem em usar build tags, alguns testes podem ficar esquecidos se a gente não se lembrar
de executalos com a tag respectiva.
Nesse caso uma outra abordagem pode ser feita utilizando variaveis de ambiente.
Exemplo abaixo:
*/

// Ao invez de build tags, controlamos quais tests devem ser executados com variavel de ambiente.
// Dessa forma quando executarmos todos os tests iremos receber uma mensagem dizendo quais tests não foram executados.
func TestInsert(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integration test")
	}

	// ...
}

/*
Uma otra opção que existe é fazer uso de .Short(). Esse metodo identifica se executamos os tests com a flag -short
e com isso podemos explicitamente deixar de executar algum teste que queremos.
Exemplo abaixo:
*/

// terminal: go test -short -v .
func TestShortFunction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping this test because I pass -short flag")
	}
}
