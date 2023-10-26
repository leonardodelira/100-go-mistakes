package codeandprojectorganization

import (
	"log"
	"net/http"
)

// Simulate variable shadowing
// A variavel client nunca é usada.
func ExampleError() {
	tracing := true
	//var client *http.Client
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return
		}
		log.Println(client)
	}
}

// Neste caso não temos erro de variable shadowing porque não estamos "redeclarando" a variavel client, mas sim atribuindo um valor.
// Trocamos := por =
func ExampleCorrect() {
	tracing := true
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
		if err != nil {
			return
		}
		log.Println(client)
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return
		}
		log.Println(client)
	}
}

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
