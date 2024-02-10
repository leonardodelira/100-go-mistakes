package errormanagement

import "fmt"

/*
Mistake 48: Panicking

Panic deve ser usado com moderação, normalmente usamos panic em duas situações.
- 1. Para sinalizar um erro claro do programador.
- 2. Quando nossa aplicação falha para criar uma dependencia obrigatória.

Como exemplo podemos ver a função "Register" do package database/sql, que usa "panic" caso não enviemos
o parametro obrigatório, que no caso é o driver.

func Register(name string, driver driver.Driver) {
    driversMu.Lock()
    defer driversMu.Unlock()
    if driver == nil {
        panic("sql: Register driver is nil")
    }
    if _, dup := drivers[name]; dup {
        panic("sql: Register called twice for driver " + name)
    }
    drivers[name] = driver
}

Outro exemplo é com a função checkWriteHeaderCode dentro de net/http que valida o código header enviado pelo programador.

func checkWriteHeaderCode(code int) {
    if code < 100 || code > 999 {
        panic(fmt.Sprintf("invalid WriteHeader code %v", code))
    }
}
*/

/*
Resumindo, panic deve ser usado com moderação. Apenas com situações onde nossa aplicação pode ser prejudicada para ser executada.
Erros de validação, chamadas externas e etc deve ser tratada de outra maneira.
*/

/*
Podemos recuperar caso um panic ocorra no fluxo da nossa aplicação.
Mas é importante lembrar que o uso de recover só funciona com conjunto com o defer
*/
func CallRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover ", r)
		}
	}()

	f()
}

func f() {
	fmt.Println("foo")
	panic("panic")
	fmt.Printf("bar")
}
