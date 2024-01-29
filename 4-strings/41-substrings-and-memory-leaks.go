package strings

import (
	"fmt"
	"strings"
)

/*
Mistake 41: Substrings and memory leaks

Ao realizar substrings podemos estar apontando para o mesmo backing array da string original, e com isso ocupar memória desnecessária.

Por exemplo: digamos que temos uma função que pega os primeiro 36 caracteres de um log, que representaria o ID desse log.
E depois salvamos esse ID em memória. Nesse momento estariamos apontando para o mesmo backing array da string original. Exemplo abaixo.
*/

type store struct {
}

func (s *store) store(uuid string) {}

func (s *store) SubstringBad(log string) {
	uuid := log[:36] //Nesse momento a variavel uuid está apontando para o mesmo backing array da variavel log. Então estaremos referenciando nossa variavel para um backing array com dados desnecessários.
	s.store(uuid)
}

func (s *store) SubstringGood(log string) {
	//Umas das soluções é forçar uma conversão para bytes para que uma nova alocação com apenas com bytes necessários (36) seja criada.
	//A partir desse momento o backing array de "uuid" será de apenas 36 bytes
	uuid := string([]byte(log[:36]))

	//Essa segunda opção está disponivel apartir do Go 1.18, a variavel clone já cria uma nova alocação em memória.
	uuid2 := strings.Clone(log[:36])

	fmt.Print(uuid, uuid2)
}
