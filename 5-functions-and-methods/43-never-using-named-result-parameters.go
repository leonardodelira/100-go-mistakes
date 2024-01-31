package functionsandmethods

/*
Mistake 43: Never using named result parameters

Usar result parameters pode ser útil para facilitar a leitura do código em determinadas condições.
Ou também pode ser útil para escrever funções pequenas.

Exemplo abaixo:
*/

// Quando vemos esse código, pode ficar dificl saber o que significa esses dois float32. Porém se usarmos result parameters a leitura fica mais fácil
func WithoutResultParameters() (float32, float32, error) {
	return 0, 0, nil
}

// Agora sabemos o que significa os retornos, e isso pode facilitar a leitura do código
func WithResultParameters() (long float32, ltd float32, err error) {
	long = 1
	ltd = 2
	return //Quando usamos result parameters apenas chamamos o "return" e o último valor atribuído as variaveis será retornado
}

/*
Resumindo, não existe uma regra para se usar result parameters. Sua aplicação nesse caso é mais para facilitar a leitura e interpretação
do código.
*/
