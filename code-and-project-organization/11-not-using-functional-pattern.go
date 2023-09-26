package codeandprojectorganization

/*
Mistake 11: Not using the functional options pattern

Esse pattern pode ser útil para criar config de uma aplicação.
Vamos imaginar uma API que precisa receber determinados valores para configurar a instancia.
Usando o functional options pattern nós podemos criar quantas configs forem necessárias, sem quebrar o client que não fornecer alguma dessas configs.
A ideia é bem parecida com um "builder"

Abaixo um exemplo de como seria o uso do código que não está comentado:

//main.go
func main() {
	config.NewServer("localhost",
		config.WithPort(8080),
		config.WithTimeout("1000"))
}

Podemos observar um pattern que é bem flexivel com o que se esperar nos parametros.
Dessa maneira nós podemos validar os parametros que foram enviados.
*/

import (
	"errors"
	"fmt"
	"net/http"
)

const defaultHTTPPort = 8080

type options struct {
	port    *int
	timeout string
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func WithTimeout(time string) Option {
	return func(options *options) error {
		options.timeout = time
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options) //Executa as funcoes que foram passadas como parametro, no caso: WithPort e WithTimeout
		if err != nil {
			return nil, err
		}
	}

	// At this stage, the options struct is built and contains the config
	// Therefore, we can implement our logic related to port configuration
	var port int
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = 9999
		} else {
			port = *options.port
		}
	}
	fmt.Println(port)
	return nil, nil
}
