package thestandardlibrary

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
Mistake 77: Common JSON-handling mistakes

77.1 - Erro ao trabalhar com JSON handling e embed types
Devemos ter em mente que quando nossa struct faz uso de embed types, nossa struct passa a ter os possíveis metodos
da interface que o tipo "embedado" tem.
Vamos a um exemplo abaixo:
*/

/*
Aqui estamos criando nossa struct embedando o time.Time, a partir desse momento nossa struct Event terá acesso a
todos os metodos de time.Time. Inclusive terá acesso ao metodo MarshalJSON() que sobrescreve o comportamento
padrão desse metodo.
Ou seja, se criarmos uma instancia de Event e realizar o Marshal o metodo que será executado é o metodo dentro de time.Time
então podemos esperar um comportamento diferente dessa ação.
Nesse caso o field "ID" não será incluido no Marshal porque o metodo MarshalJSON() que existe dentro de time.Time obviamente
não conhece esse campo.
Resumindo: quando fazemos uso de embedded values, os metodos do valor embadado terá prioridade na execução. Tem um comportamento
de override.
*/
type Event struct {
	ID string
	time.Time
}

func Mistake77_EmbbedJsonWrong() {
	event := Event{
		ID:   "1234",
		Time: time.Time{},
	}

	/*
		É nesse momento que o ID é ignorado, o metodo de MarshalJSON que está sendo executado é o que pertence ao time.Time
	*/
	result, _ := json.Marshal(event)

	fmt.Print(string(result))
}

/*
Para resolver o problema acima, basta não fazer uso de embbeded value ou
criar a propria implementação de MarshalJSON() para a struct Event.
*/
