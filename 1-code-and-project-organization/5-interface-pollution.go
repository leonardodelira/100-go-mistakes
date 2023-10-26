package codeandprojectorganization

/*
#5: Interface pollution
Este é um erro comum de programadores que vem de outros linguagens de programação com certos habitos.

"As abstraçoes devem ser descobertas, não criadas". O que isso significa? Nós não devemos começar a criar abstraçoes
em nosso código se não temos razão imediata para isso.
	Nós devemos criar uma interface quando nós precisarmos dela e não quando prevemos que vamos precisar dela.


Tres casos que podem nos dar uma ideia geral de quando usar interface:
-Common behavior
--
** Decoupling **
No exemplo abaixo vemos um alto acoplamento na struct, isso pode dificultar nossos testes unitários. Uma maneira
de resolver isso é por usar interface. Esse é um bom cenário para usar interface.

type CustomerService struct {
    store mysql.Store          ❶
}

func (cs CustomerService) CreateNewCustomer(id string) error {
    customer := Customer{id: id}
    return cs.store.StoreCustomer(customer)
}

Mesmo exemplo porém com interfaces:

type customerStorer interface {      ❶
    StoreCustomer(Customer) error
}

type CustomerService struct {
    storer customerStorer            ❷
}

func (cs CustomerService) CreateNewCustomer(id string) error {
    customer := Customer{id: id}
    return cs.storer.StoreCustomer(customer)
}

Neste caso o uso da interface é válido porque nos dá mais flexibilidade em como vamos testar o metodo.

** Restricting behavior **
Usar interface pode ser útil em Go quando queremos restringir que um type tenha comportamento especifico.
Por exemplo, imaginos que tenhamos o seguinte arquivo de configuração dinamica:

type IntConfig struct {
    // ...
}

func (c *IntConfig) Get() int {
    // Retrieve configuration
}

func (c *IntConfig) Set(value int) {
    // Update configuration
}

Quando formos fazer uso dessa config em nosso código, podemos forçar que o type que fará uso dessa config não poderá
editar nenhum dado, apenas fazer o "get" do value. Para isso podemos usar uma interface e colocar ela como
injeção de dependencia no nosso type.

type intConfigGetter interface {
    Get() int
}

type Foo struct {
    threshold intConfigGetter
}

func NewFoo(threshold intConfigGetter) Foo {    ❶
    return Foo{threshold: threshold}
}

func (f Foo) Bar()  {
    threshold := f.threshold.Get()              ❷
    // ...
}

Na configuração acima nós criamos uma interface que tem apenas o metodo Get, que por meio implicito vai satisfazer o struct
IntConfig e dessa maneira vamos ter acesso apenas ao GET e não será possível editar nenhum valor.
*/

type IntConfig struct {
	// ...
}

func (c *IntConfig) Get() int {
	return 0
}

func (c *IntConfig) Set(value int) {
	// Update configuration
}

type initConfigGetter interface {
	Get() int
}

type Foo struct {
	config initConfigGetter
}

func NewFoo(initConfig initConfigGetter) Foo {
	return Foo{config: initConfig}
}

func (f Foo) Bar() {
	f.config.Get() //Aqui podemos acessar o GET do struct IntConfig (Se ele for injetado é claro)
}
