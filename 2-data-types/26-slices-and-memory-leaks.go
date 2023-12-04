package datatypes

/*
Mistake 26: Slices and memory leaks

Duas categorias de "memory leak"
1. capacity
2. pointers

*/

/*
1. Capacity:
Imagine que nós estamos lendo mensagens de um brokers, e cada mensagem contem 1000 bytes.
Os cinco primeiros bytes de slice contem o "type" da nossa mensagem.
Agora vamos considerar que em determinado momento vamos precisar recuperar esse "type" e salvar em memória.
O que seria uma má pratica nesse contexxto? O slicing mal feito. Exemplo abaixo:
*/

func consumeMessages() {
	for {
		msg := receiveMessage()
		// Do something with msg
		storeMessageType(getMessageType(msg))
	}
}

/*
O slice realizado na função abaixo tem um problema, ele retorna um novo slice de bytes porém mantem o capacity do slice original Que no caso seria de 1000 bytes.
Com isso nós mantemos em memória valores desnecessários e nossa aplicação irá consumir muita memória.
*/
func getMessageType(msg []byte) []byte {
	return msg[:5] //faz o slicing mas mantem o capacity do slice original, se o slice original for muito grande teremos problemas de memória.
}

func receiveMessage() []byte       { return []byte("a") }
func storeMessageType(byte []byte) {}

// Uma maneira de corrigir o problema é fazer o slicing e também definir qual será o capacity desse novo slice, dessa maneira não mantemos
// espaço reservado no backng array desnecessário
func getMssageTypeCorrect(msg []byte) []byte {
	return msg[:5:5]
}

/*
2. Pointers:
O mesmo problema acima pode acontecer com um slice de structs.
Se temos um slice de structs e essas structs possuem alguma property que gerará um ponteiro (exemplo um array ou outro slice). Caso a gente faça
um slicing desses valores (ex: foos[:3]) os structs que sobraram desse slicing continuará em memoria.
O ideial é sempre fazer um slicing e redefinir o capacity de nossa struct, ex: foos[:3:3].
	Ou até mesmo usar a função copy()
*/
