package functionsandmethods

import (
	"errors"
	"strings"
)

/*
Mistake 45: Returning a nil receiver

"nil" é um "zero value" de: pointer, channel, func, interface, map, or slice type.
ou seja, se nossa função retorna um ponteiro para nil, quem está chamando essa função pode ser enganado.

Por exemplo o metodo Validate abaixo, ele retorna um error.
Se observarmos a implementação do method, nós sempre retornamos o valor de "m". O que acontece nesse caso é
que "m" inicia com um ponteiro para nil e depois retornamos esse valor.

Quem for chamar esse metodo irá receber como se err != nil, porque na verdade está recebendo um ponteiro.
Mas quando vamos ver o valor que temos no ponteiro, encontramos "nil".

O ideal nesse caso é de forma explicita retornar "nil" na nossa função. Dessa forma o "caller" não irá ser enganado
recebendo um ponteiro que aponta para nil.
*/

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Name string
	Age  int
}

func (c Customer) Validate() error {
	var m *MultiError

	if c.Name == "" {
		m = &MultiError{}
		m.Add(errors.New("name cant be empty"))
	}

	if c.Age < 0 {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("age invalid"))
	}

	return m
}
