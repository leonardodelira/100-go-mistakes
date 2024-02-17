package concurrencypractice

import (
	"fmt"
	"sync"
)

/*
Mistake 68: Forgetting about possible side effects with string formatting


*/

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

/*
Considerando que estamos sobrescrevendo o metodo String(). Temos que ficar atentos para não sofrermos com deadlock.
No exemplo abaixo nós enviamos um error caso o "age" seja menor que 0.
Porém como estamos sobrescrevendo o metodo String e criando um lock da nossa struct, nesse caso teremos um deadlock.
Quando o metodo String() ser chamado ele não vai conseguir ler o dado de Customer porque um lock já está ativo.
*/
func (c *Customer) UpdateAge(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}

	c.age = age
	return nil
}

func (c *Customer) ShowAge() int {
	return c.age
}

func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}
