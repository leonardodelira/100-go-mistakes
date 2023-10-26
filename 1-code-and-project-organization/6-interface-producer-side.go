package codeandprojectorganization

/*
Definição de "Producer side"
	Producer side—An interface defined in the same package as the concrete implementation

Definição de "Consumer side"
	Consumer side—An interface defined in an external package where it’s used

Em GO, na maioria das vezes devemos sempre lembrar que abstrações devem ser "descobertas". O que isso significa?

Digamos que tenhamos a seguinte interface e sua implementação

package store

type CustomerStorage interface {
    StoreCustomer(customer Customer) error
    GetCustomer(id string) (Customer, error)
    UpdateCustomer(customer Customer) error
    GetAllCustomers() ([]Customer, error)
    GetCustomersWithoutContract() ([]Customer, error)
    GetCustomersWithNegativeBalance() ([]Customer, error)
}

O que fazer caso queira usar algum dos metodos dessa interface em outro package, por exxemplo "client".
Neste outro package (neste caso client), podemos criar uma interface que irá atender apenas a necessidade que temos, por exemplo:

package client

type customersGetter interface {
    GetAllCustomers() ([]store.Customer, error)
}

type clientSomeStuff struct {
	customer customersGetter
}

Podemos criar uma struct e criar uma dependencia que tenha o tipo dessa nova interface "customersGetter"
Quando formos injetar o tipo concreto da implementação, teremos acesso apenas ao método que necessitamos em nosso contexto.
*/
