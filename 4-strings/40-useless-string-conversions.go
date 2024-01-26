package strings

import (
	"bytes"
	"io"
	"strings"
)

/*
Mistake 40: Useless string conversions

As maiorias das operações de I/O é feita com []bytes.
Muitas vezes preferimos trabalhar com string pela conveniencia, porém é importante considerar que trabalhar com []bytes pode evitar
conversões desnecessárias já que o pacote bytes oferece muitas das mesmas funções do pacote string.

Sempre que possível podemos tentar manipular os bytes sem fazer a conversão para strings.
*/

/*
Por exemplo nesse caso abaixo, estamos convertendo o []byte para string, realizamos o "trim" e depois temos que
converter para bytes novamente porque nossa função retorna []byte.
Isso pode ser evitado, porque como dito o package byte possuí várias funções iguais ao package string
*/
func getBytes(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return []byte(sanitize(string(b))), nil
}

func sanitize(s string) string {
	return strings.TrimSpace(s)
}

// Forma correta sem conversão desnecessárias
func getBytes2(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return sanitize2(b), nil
}

// Go tambem consegue remover espaçoes vazios dos bytes, entre outras funções.
func sanitize2(s []byte) []byte {
	return bytes.TrimSpace(s)
}
