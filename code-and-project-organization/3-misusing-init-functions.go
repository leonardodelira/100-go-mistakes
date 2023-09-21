package codeandprojectorganization

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

/*
init function é a primeira função a ser executada em um package.
devemos ficar atentos porque ela tem algumas limitações, como por exemplo, não retorna valores.

ps: os imports de um package são a primeira coisa a ser executada,
e se esses imports tem funcoes init, elas serão executadas antes de qualquer outra coisa.
*/

var db *sql.DB

/*
Aqui temos um mal exemplo de como utilizar init func,
- Talvez não iriamos querer que um erro na conexão com o banco de dados gera-se um panic. Poderiamos querer implementar um retry.
- Outro ponto é que em testes unitários nós teriamos problemas com a init function, não é possível testa-la de forma isolada. Neste caso abaixo nós criariamos uma conexão com o banco desnecessária.
*/
func init() {
	dataSourceName :=
		os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	db = d
}

/*
Por essas razoes, a função de inicializacão deveria ser tratada como uma func qualquer
*/
func createClient(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

/*
Mas em que casos pode ser vantajoso fazer uso da init func?
For example, the official Go blog (http://mng.bz/PW6w) uses an init function to set up the static HTTP configuration:

Neste exemplo nós não temos gerenciamento de erros
Nenhuma variavel global
E a função não irá impactar os tests unitários.
*/
func init() {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)

	http.Handle("/lib/godoc/", http.StripPrefix("/lib/godoc/", static))
}
