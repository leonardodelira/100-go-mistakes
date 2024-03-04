package thestandardlibrary

import (
	"net"
	"net/http"
	"time"
)

/*
Mistake 81: Using the default HTTP client and server

Ao usarmos o http client default do Go, podemos esquecer algumas configurações importantes para garantir um bom
desempenho de nossa aplicação.

As configurações são:
- net.Dialer.Timeout: Specifies the maximum amount of time a dial will wait for a connection to complete.
- http.Transport.TLSHandShakeTimeout: Specifies the maximum amount of time to wait for the TLS handshake.
- http.Transport.ResponseHeaderTimeout: Specifies the amount of time to wait for a server’s response headers.
- http.Client.Timeout: Specifies the time limit for a request. It includes all the steps, from step 1 (dial) to step 5 (read the response body).

Para sobrescrever essas configurações, podemos seguir o exemplo abaixo.
*/

/*
Na função abaixo estamos sobrescrevendo também o nosso pool de conexões. (IdleConnTimeout)
Por padrão o uma conexão fica no pool por apenas 90 segundos, mas podemos aumentar esse número como está feito na linha 37.

Porém é importante também alterar o MaxIdleConnsPerHost, esse parametro define quantas idle connecitons ficará aberta por host,
por default o valor é 2, isso significa que se tivermos 100 requests concorrentes apenas 2 serão aproveitadas do pool de conexões
e o restante terá que ser criada novamente, isso é custoso. Nesse caso podemos parametrizar um valor maior para o pool. Linha 42.
*/
func OverrideClientHTTP() {
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   time.Second,
			ResponseHeaderTimeout: time.Second,
			IdleConnTimeout:       120 * time.Second,
			MaxIdleConnsPerHost:   100,
			MaxIdleConns:          100,
		},
	}

	client.Get("...")
}

type HttpHandler struct{}

func (h *HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}

/*
Temos que ficar atentos ao trabalhar com o server default do golang também.
Algumas configurações devem ser sobreescritas para evitar por exemplo uma chamada http que vá durar muito tempo e irá consumir
muitos recursos de nosso servidor.

- http.Server.ReadHeaderTimeout: A field that specifies the maximum amount of time to read the request headers
- http.Server.ReadTimeout: A field that specifies the maximum amount of time to read the entire request
- http.TimeoutHandler: A wrapper function that specifies the maximum amount of time for a handler to complete

Por exemplo se o ReadTimeout não estiver setado, nossa aplicação não terá timeout e a conexão se manterá aberta
até que o client feche a conexão.
*/
func OverrideServerHTTP() {
	handler := &HttpHandler{}
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,
		Handler:           http.TimeoutHandler(handler, time.Second, "foo"),
	}

	server.ListenAndServe()
}
