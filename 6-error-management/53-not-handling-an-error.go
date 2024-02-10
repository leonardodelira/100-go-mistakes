package errormanagement

import "errors"

/*
Mistake 53: Not handling an error

Podemos ignorar erros no Golang, porém devemos fazer isso de forma explicita para deixa claro para futuros
leitores de nosso código que ignoramos o retorno de erro de forma intencional
*/

func IgnoreErrorCorrect() {
	_ = notify()
}

/*
Ignorar que a função retorna um erro também funciona. Mas isso pode atrapalhar futuros leitores do código.
Será que ignoramos o erro de forma intencional?
Será que esquecemos de tratar esse erro?
Então nesse caso, sempre devemos deixar explicito que ignoramos o erro que a função retorna com um "blank identifier"
*/
func IgnoreErrorIncorrect() {
	notify()
}

func notify() error {
	return errors.New("some error")
}
