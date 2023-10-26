package codeandprojectorganization

/*
Mistake 14: Ignoring package name collisions

Nomes de pacotes podem colidir com nomes de variaveis, fazendo com que os pacotes não possa ser reutilizados no escopo.
Exemplo:
``
package redis

type Client struct { ... }

func NewClient() *Client { ... }

func (c *Client) Get(key string) (string, error) { ... }
``

O pacote redis tem como nome "redis", como vamos utilizar do lado do client normalmente vamos criar nossa váriavel de "redis" também.

``
package internal

redis := redis.NewClient()
v, err := redis.Get("foo")
``

Nesse caso acima, a referencia para a lib é perdida dentro do escopo da variavel "redis".

Para evitar isso podemos obviamente alterar o nome da variavel ou trabalhar com import alias.

``
import redisapi "mylib/redis"

// ...

redis := redisapi.NewClient()
v, err := redis.Get("foo")
``

Agora o import do pacote tem outro nome e podemos referenciar ele a qualquer momento.
*/
