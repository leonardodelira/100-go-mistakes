package errormanagement

import (
	"errors"
	"fmt"
)

/*
Mistake 49: Ignoring when to wrap an error

É interessante usar error wrapping em geral em duas situações.
Quando queremos:
- Adicionar mais contexto ao nosso erro.
- Marcar o erro como um erro especifico.
*/

type ErrorBar struct {
	Err error
}

func (b ErrorBar) Error() string {
	return "bar failed: " + b.Err.Error()
}

/*
Para adicionar mais contexto ao nosso erro, basta utilizarmos Errorf.
*/
func AddContextError() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("some error in foo: %w", err)
	}
	return nil
}

/*
Porém caso a gente queira tipar o erro, também é possível.
Podemos criar a struct e fazer com que ela respeite a interface error.
No caso criamos a ErrorBar
*/
func AddTypeError() error {
	err := bar()
	if err != nil {
		return err
	}
	return nil
}

func bar() error {
	return ErrorBar{Err: errors.New("some error on bar")}
}
