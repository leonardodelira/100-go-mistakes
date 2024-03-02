package thestandardlibrary

import "net/http"

/*
Mistake 80: Forgetting the return statement after replying to an HTTP request

Lembrar que http.Error não para a execução da função.
Então temos que colocar "return" caso algum erro ocorra.
*/

func handler(w http.ResponseWriter, req *http.Request) {
	err := foo(req)
	if err != nil {
		http.Error(w, "foo", http.StatusInternalServerError)
		return //Como dito, http.Error não para o processo da função então temos que usar o return
	}

	// ...
}

func foo(req *http.Request) error {
	return nil
}
