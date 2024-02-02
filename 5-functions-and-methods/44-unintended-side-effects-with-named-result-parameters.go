package functionsandmethods

import "context"

/*
Mistake 44: Unintended side effects with named result parameters

Devemos lembrar que quando fazemos uso de named result parameters, os parametros que criamos no retorno
são iniciados com valor zero respectivo ao seu tipo.
Enão dependendo da nossa implementação, não podemos esquecer de atribuir o valor que queremos retornar nesses
named result parameters.
Exemplo abaixo:
*/

// Aqui todas as variaveis foram iniciadas com seu respectivo valor 0, incluse o err com nil
func nameResultParametersSideEffect(ctx context.Context, address string) (latitude float32, longitude float32, err error) {
	//some validates
	//some code

	/*
		Aqui está o side effect que podemos sofrer, estamos retornando "err" sem atribuir um valor para ele. Devemos lembrar que ele foi apenas iniciado.
		Ou seja, nesse caso estaremos retornando nil (err=nil)
		O correto seria antes do retorno, atribuir o valor a ele
		err = ctx.Err()
	*/
	if ctx.Err() != nil {
		return 0, 0, err
	}
	return
}

/*
Um outro ponto de atenção, quando fazemos uso de name result parameters não necessáriamente nós precisamos fazer uso de naked return, ou seja
apenas colocar o "return". Ainda podemos retornar os valores que queremos como mostra no exemplo da linha 26.
*/
