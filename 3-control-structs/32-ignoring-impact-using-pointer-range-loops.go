package controlstructs

import "fmt"

/*
Mistake 32: Ignoring the impact of using pointer elements in range loops

Ao fazermos um range loop em pointers, podemos referenciar os elementos errados se não estivermos atentos.
Devemos lembrar que ao fazermos um range loop, uma variavel nova é criada. E no caso essa variavel tem apenas um endereço de memória.
Então abaixo vemos o erro.
*/

type Customer struct {
	ID      int
	Balance float64
}

type Store struct {
	m map[int]*Customer
}

func (s *Store) ExampleMistake32(customers []Customer) {
	s.m = make(map[int]*Customer)
	for _, customer := range customers {
		s.m[customer.ID] = &customer
	}
	/*
		O valor dessa variavel será sempre o último elemento do slice de customer, porque?
		Quando fazemos o range loop, uma "pre variavel" foi criada com apenas um endereço de memória. (No caso "customer")
		E como estamos atribuindo o endereço de memória ao nosso Store, ele sempre terá o endereço de memória da "pre variavel"
	*/
	for _, c := range s.m {
		fmt.Println(c.ID, c.Balance) //Todos os valores serão iguais ao ultimo elemento do slice customers que recebemos como input.
	}
}

// Para corrigir a situação, podemos declarar uma variavel dentro do nosso range loop, dessa forma a cada iteração uma nova variavel será criada
// e por conta disso um novo endereço de memória. Então a partir de agora os valores corretos serão atribuídos ao nosso map.
// Porque não estamos mais referenciando a "pre variavel" do range loop, mas sim a nova variavel criada dentro do loop.
func (s *Store) ExampleCorrect32(customers []Customer) {
	s.m = make(map[int]*Customer)
	for _, customer := range customers {
		current := customer
		s.m[current.ID] = &current
	}

	for _, c := range s.m {
		fmt.Println(c.ID, c.Balance) //Agora os valores estão corretos, porque criamos uma nova variavel dentro do loop e por conta disso tivemos um novo endereço de memoria por variavel
	}
}
