package codeandprojectorganization

import "sync"

/*
Mistake 10: Not being aware of the possible problems with type embedding

Quando fazemos o embed de uma struct na outra, ela passa a ter acesso aos seus fields (o que é obvio).
Porém devemos ficar atentos a cenários onde não queremos que clients externos tenham acesso aos fields de nossa field.
*/

// Tendo acesso a todos fields do embbed
type Bar struct {
	Baz int
}

type Foos struct {
	Bar
}

func Example() {
	foos := &Foos{}
	foos.Baz = 10
	foos.Bar.Baz = 12
}

// Aqui também temos acesso a todos fields da struct, porém pode ser que nesse cenário não era isso que queriamos
type Client struct {
	sync.Mutex
	ID int
}

/*
Aqui nós estamos retornando a struct Client toda, qualquer client que fizer o "new" dessa struct
terá acesso os metodos de Mutex, porque ele foi embedado e seus metodos são publics.
*/
func (c *Client) New() *Client {
	c.Lock()
	c.ID = 10
	c.Unlock()
	return &Client{}
}

// Caso a ideia seja que clientes externos não tenha acesso a determinas fields, devemos nomealas.
type ClientV2 struct {
	mu sync.Mutex //nomeando o tipo para que ele não seja "embedded"
	ID int
}

// Agora nesse cenário clientes externos não conseguem ter acesso a "mu" porque ele é privado.
func NewV2() *ClientV2 {
	return &ClientV2{}
}

func (cv2 *ClientV2) SetID() {
	cv2.mu.Lock()
	cv2.ID = 10
	cv2.mu.Unlock()
}
