package thestandardlibrary

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

/*
Mistake 78: Common SQL mistakes
*/

/*
sql.Open não necessáriamente estabelece uma conexão com o banco de dados. Ele apenas valida se os argumentos
são validos. Porém isso por driver.
No caso do postgres, para garantir a conexão devemos fazer uso do .Ping()
*/
func Mistake78_SqlOpen() error {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=dbtest sslmode=disable")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Print("Ping Success")
	return nil
}

/*
Ao criarmos uma conexão com o banco, não podemos esquecer que na verdade estamos criando um pool de conexões.
As conexões nesse pool podem obter os seguintes status: "Already used" e "Idle".

Already used = está sendo usada
Idle = está criada e disponível para uso.

Além disso, podemos configurar o nosso pool de conexões com as seguintes propriedades.
- SetMaxOpenConns = número máximo de conexões ativas com o banco
- SetMaxIdleConns = número máximo de conexões criadas mas fora de uso
- SetConnMaxIdleTime = tempo máximo que uma conexão irá existir no stado de idle
- SetConnMaxLifetime = tempo máximo que uma conexão pode ser mantida aberta antes de fechar
*/

/*
Outro fator é não usar o "preperad statment"
É importante usar o "preparad stament" para consultas SQLs que serão executadas repetidas vezes.
Isso porque internamente o SQL precompila essa query e com isso ganhamos em efeciencia e segurança. Nas próximas vezes
que executarmos essa query ela já vai estar compilada.
Abaixo exemplo de prepared statment
*/
func Mistake78_PreparedStatment() error {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=dbtest sslmode=disable")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	//Importante fazer uso de Prepare para deixar esse consulta compilada internamente no SQL.
	//E também por questão de segurança por evitar SQL Injections
	stmt, err := db.Prepare("SELECT * FROM ORDER WHERE ID = ?")
	if err != nil {
		return err
	}
	id := 1
	rows, _ := stmt.Query(id)
	for rows.Next() {
		///...
	}
	return nil
}

/*
Outro ponto de atenção que devemos ter com nossas consultas SQL é ao lidar com null values.
Digamos que determinado campo do banco de dados retorne null, como devemos tratar isso ao realizar o
row.Scan de nossos valores?
Primeiro vamos ver um exemplo do que pode acontecer.
*/
func Mistake78_NullValues() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=dbtest sslmode=disable")
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	/*
		Digamos que o valor de "DEP" é null na nossa base de dados.
		Nesse vamos ao realizar o rows.Scan na linha 113 teremos um erro. Porque nossa variavel
		não está esperando um valor nulo.

		Nesse caso podemos ter duas soluções, criar a variavel do tipo "ponteiro para string"
		ex: departament *string
		ou
		Usar o tipo nullable do nosso driver
		ex: departament sql.NullString
		A partir desse momento já podemos receber valores nulos da nossa base de dados sem ter erro.
	*/
	id := 1
	rows, err := db.Query("SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return
	}
	defer rows.Close()

	var (
		department string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return
		}
		// ...
	}
}

/*
Devemos lembrar de validar algum possível erro no rows.Next()
Para fazer isso basta considerar o que está na linha 172
*/
func Mistake78_RowInteractionErrors() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=dbtest sslmode=disable")
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	fmt.Print("Ping Success\n")

	id := 1
	rows, err := db.Query("SELECT dep, age FROM emp WHERE id = $1", id)
	if err != nil {
		fmt.Printf("error on query: %v", err)
		return
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var (
		department sql.NullString
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			fmt.Printf("failed to scan rows: %v\n", err)
			return
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("failed to rows next: %v\n", err)
		return
	}

	fmt.Println(department.Value())
	fmt.Println(age)
}
