package concurrencypractice

import (
	"context"
	"net/http"
	"time"
)

/*
Mistake 61: Propagating an inappropriate context

Devemos tomar cuidado ao propaar nosso contexto no fluxo da nossa aplicação, ainda mais quando trabalhamos com processos
assincronos.
Por exemplo, digamos que temos um handler que faz alguma ação e antes de enviar o response para o client nós queremos enviar
uma mensagem no kafka de forma assincrona para não influenciar o tempo de resposta do nosso handler.
Digamos que a função que enviar a mensagem para o kafka espera como parametro um context, o que faremos nesse caso? Compartilhamos
o mesmo context da request?
	A resposta é não, porque se escrevermos a resposta ao cliente antes do processo assincrono enviar a mensagem ao kafka. Nós
	teremos um erro. Porque? Por que o contexto compartilhado com o processo assincrono vai receber o sinal de "cancel" visto que
	processo do handler já terminou.
*/

func doSomeTask(c context.Context, r *http.Request) (any, error) {
	return nil, nil
}

func writeResponse(w http.ResponseWriter, response any) {
	w.Write([]byte("success"))
}

func publish(c context.Context, value any) {

}

/*
Como dito, aqui temos um mal exemplo de como propagar o contexto da request.
A função writeResponse pode ser executada primeiro, antes da goroutine terminar o processo. Dessa forma a função publish
que está recebendo o mesmo contexto da requisição irá falhar, porque irá receber o evento de "cancel()"
*/
func BadExamplePropagatingContext(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	go func() {
		publish(r.Context(), response)
	}()

	writeResponse(w, response)
}

/*
Uma primeira solução que podemos usar é por enviar um context novo para a função publish, sem vinculo com o contexto da
request.
Mas e se o contexto da request possui valores que queremos manter no contexto da função publish? Nesse caso podemos criar
um contexto personalizado apenas com o valor do contexto original. Exemplo abaixo.
*/
func GoodExamplePropagatingContext(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	go func() {
		publish(context.Background(), response)
	}()

	writeResponse(w, response)
}

/*
Criando um contexto personalizado para manter o valor do contexto original mas sem ter vinculo com os eventos disparados pelo
original.

Basta criarmos uma struct que recebe um ctx e que respeite a assinatura da interface de um contexto.
As funções de disparo de erro do contexto original (deadline, done, err) foram sobescrita para não disparar nenhum erro.
Apenas aproveitamos o valor "Value" do contexto original e passamos para a função publish.
*/

type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
	return nil
}

func (d detach) Err() error {
	return nil
}

func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}

/*
Agora nosso handler pode finalizar o processo e não teremos problemas em nossa goroutine.
Ela está executando de forma independente (como se espera) e herdando o valor do context da request.
*/
func BetterExamplePropagatingContext(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	go func() {
		publish(&detach{ctx: r.Context()}, response)
	}()

	writeResponse(w, response)
}
