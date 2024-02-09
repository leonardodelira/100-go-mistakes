package functionsandmethods

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

/*
Mistake 46: Using a filename as a function input

Passa o nome do arquivo para nossa função que irá ler o arquivo é uma má pratica em muitos casos.
Especialmente quando precisamos criar os units tests.
*/

/*
Exemplo dessa função abaixo que conta o número de linhas vazias em um arquivo.
Quando formos testar ela, vamos ter que criar vários arquivos diferentes para testar diferentes casos de uso.
Então quanto mais complexo for nossa função, mais arquivos diferentes com os casos de uso deveriamos criar.

Além disso a função não é reutilizavel, se precisarmos por exemplo contar o número de linhas vazias de uma HTTP request,
teriamos que criar outra função.

	func countEmptyLinesInHTTPRequest(request http.Request) (int, error) {
	    scanner := bufio.NewScanner(request.Body)
	    // Copy the same logic
	}
*/
func countEmptyLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// Handle file closure

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}

	return -1, nil
}

/*
Uma forma de evitar esses problemas é por passar como parametro a função um io.Reader.
Agora nossa função fica mais abstrata para diferente inputs que satisfazem a interface io.Reader.
Pode ser um arquivo, http request, socket ...
Sem contar que para criar os testes fica mais simples também.
*/
func countEmptyLines(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// ...
	}
	return -1, nil
}

/*
Exemplo de como o test fica mais simples.
Não precisamos criar arquivos para gerar os cenários.
*/
func TestCountEmptyLines(t *testing.T) {
	emptyLines, err := countEmptyLines(strings.NewReader(
		`foo
            bar
 
            baz
            `))
	fmt.Print(emptyLines, err)
}
