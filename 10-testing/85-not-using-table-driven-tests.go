package testing

import "strings"

/*
Mistake 85: Not using table-driven tests

Usar table-driven testes nos ajuda a evitar duplicação de código.
Verificar exemplo de uso no arquivo _test.go
*/

func removeNewLineSuffixes(s string) string {
	if s == "" {
		return s
	}
	if strings.HasSuffix(s, "\r\n") {
		return removeNewLineSuffixes(s[:len(s)-2])
	}
	if strings.HasSuffix(s, "\n") {
		return removeNewLineSuffixes(s[:len(s)-1])
	}
	return s
}
