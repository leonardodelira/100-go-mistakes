package errormanagement

import (
	"errors"
	"fmt"
)

/*
Mistake 50: Checking an error type inaccurately

Podemos verificar o tipo de erro que recebemos para fazer alguma tratativa especial.
Porém temos que ficar atentos a forma que validamos o tipo de erro que recebemos.
*/

type ErrorTransactionDB struct {
	Err error
}

func (e ErrorTransactionDB) Error() string {
	return "some error on DB: " + e.Err.Error()
}

/*
Para validar o tipo de erro que temos, podemos fazer uso da função errors.As. Foi introduzida a partir do Go 1.13.

Outra maneira de fazer a validação é por usar um swith case con err(.type), mas essa forma está defasada porque pode ocorrer
do nosso error estar "wrapped" dentro de outro erro, e dessa forma o condicional "swith case con err(.type)" não funciona.

A vantagem do errors.As é que ele é recursivo até chegar no ultimo nível de "wrapper"
*/
func HandlerMock(transactionID string) {
	_, err := getTransactionAmout("12345")
	if err != nil {
		if errors.As(err, &ErrorTransactionDB{}) {
			fmt.Println("tivemos um erro de db")
			return
		}
		fmt.Printf("tivemos qualquer outro tipo de erro que não foi tratado")
	}
}

/*
Na primeira validação apenas retornamos um tipo "Erro", padrão do golang.
Porém a função getAmountFromDB retorna um erro do tipo ErrorTransactionDB. Então podemos recuperar o tipo de erro no "caller"
Neste caso temos que fazer uso do %w para manter o tipo do erro para o caller, ou simplesmente retornar err.

Se fizermos uso de %v, o caller não irá receber que o tipo de error é ErrorTransactionDB.
*/
func getTransactionAmout(transactionID string) (float32, error) {
	if len(transactionID) < 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getAmountFromDB(transactionID)
	if err != nil {
		return 0, fmt.Errorf("some error on transaction amount: %w", err)
	}

	return amount, nil
}

func getAmountFromDB(transactionID string) (float32, error) {
	return 0, ErrorTransactionDB{Err: errors.New("database some error")}
}
