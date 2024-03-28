package testing

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
Mistake 88: Not using testing utility packages

httptest é um pacote muito útil que pode nos ajudar a testar http server e client. Dessa maneira nao precisamos reinventar a roda.
*/

func Mistake88Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost",
		strings.NewReader("foo"))
	w := httptest.NewRecorder()
	Mistake88Handler(w, req)

	if got := w.Result().Header.Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("api version: expected 1.0, got %s", got)
	}

	body, _ := ioutil.ReadAll(w.Body)
	if got := string(body); got != "hello foo" {
		t.Errorf("body: expected hello foo, got %s", got)
	}

	if http.StatusOK != w.Result().StatusCode {
		t.FailNow()
	}
}
