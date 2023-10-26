package codeandprojectorganization

/*
Mistake 13: Creating utility packages

criar um package chamado utils, commoms, shared e etc não é uma boa pratica
nomear os pacotes dessa maneira não diz nada a respeito do que eles fazem.

Dentro de utils podemos ter qualquer coisa, isso não ajuda na hora de ler o código.
Exemplo ruim :
//util.go
package util

func NewStringSet(...string) map[string]struct{} {    ❶
    // ...
}

func SortStringSet(map[string]struct{}) []string {    ❷
    // ...
}

//client.go
set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))

=========================================================

Exemplo bom:
//stringset.go
package stringset

func New(...string) map[string]struct{} { ... }
func Sort(map[string]struct{}) []string { ... }

client.go
set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))

Neste segundo exemplo, criamos um pacote com uma definição clara do que ele faz.
E a maneira de fazer uso dele também é mais elegante.
*/
