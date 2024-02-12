package errormanagement

import (
	"database/sql"
	"log"
)

/*
Mistake 54: Not handling defer errors

Como sabemos, podemos fazer uso de defer para uma funcão ser executado no fim do processo de outra função.
Porém pode ocorrer de que a função que iremos executar com defer tenha algum erro, e por isso precisamos
validar corretamente.
*/

const query = "..."

// Bad example
func getBalance(db *sql.DB, clientID string) (
	float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// Use rows

	return 0, nil
}

// Good Example
func getBalance2(db *sql.DB, clientID string) (
	float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	/*
		Como vimos, se vamos ignorar algum erro que a função retorna, devemos deixar isso de forma explicita.
		Mesmo se chamamos alguma função com defer.
	*/
	defer func() {
		_ = rows.Close()
	}()

	// Use rows

	return 0, nil
}

// Other Good Example
func getBalance3(db *sql.DB, clientID string) (
	balance float32, err error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	/*
		Mas e se quisermos retornar os erros que rows.Close() pode retornar dentro de um defer?
		Devemos fazer uso de named result params. Agora o "caller" irá receber o erro
	*/
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	// Use rows

	return 0, nil
}
