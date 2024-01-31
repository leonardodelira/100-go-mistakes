package functionsandmethods

import "fmt"

/*
Mistake 42: Not knowing which type of receiver to use

Temos dois tipos de receiver: values e pointers.
Cada um deles tem uma proposta diferente.

Quando fazemos uso de receiver value, nós basicamente estamos criando uma variavel local. Dessa forma qualquer alteração feita no receiver
terá efeito apenas localemente no method.
Exemplo abaixo
*/

type customer struct {
	balance float64
}

func (c customer) ExampleReceiverValue(v float64) {
	c.balance += v
	fmt.Print("balance ", c.balance) //result 150
	fmt.Println()
}

func TestReceiverValue() {
	customer := &customer{
		balance: 100,
	}
	customer.ExampleReceiverValue(50)
	fmt.Print("customer after add value ", customer.balance)
	//result 100, o valor adicionado não será exibido para nosso method "criou uma variavel local"
}

type customer2 struct {
	balance float64
}

// Ao usarmos o ponteiro para customer2, nós iremos alterar a struct original sem criar uma cópia local.
func (c *customer2) ExampleReceiverPointer(v float64) {
	c.balance += v
	fmt.Print("balance ", c.balance) //result 150
	fmt.Println()
}

func TestReceiverPointer() {
	customer := &customer2{
		balance: 100,
	}
	customer.ExampleReceiverPointer(50)
	fmt.Print("customer after add value ", customer.balance)
	//result 150, nosso method está recebendo o ponteiro da struct customer. Dessa forma ele irá alterar o valor da struct original e não irá
	//crir uma cópia local.
}
