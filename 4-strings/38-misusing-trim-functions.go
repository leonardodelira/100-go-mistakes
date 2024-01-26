package strings

/*
Mistake 38: Misusing trim functions

As funções abaixo tem diferentes resultados:
TrimRight
TrimLeft
TrimPrefix
TrimSuffix

Devemos ficar atentos e usa-las corretamente.

Essas funções abaixo, como o nome já diz .. começa pela direta ou pela esqueda.
O que elas fazem é verificar se na string tem os elementos que queremos remover, no caso "xo" ou "ox"
Essa verificação é feita elemento por elemento, e não ele como um todo .. ou seja, primeiro ve se tem "o" e depois "x".
fmt.Println(strings.TrimRight("123oxo", "xo")) //123
fmt.Println(strings.TrimLeft("oxo123", "ox")) // 123

======

Já essas funções abaixo fazem o trim com base em todo o input e não elemento por elemento.
Por exemplo, "xo" e "ox".
fmt.Println(strings.TrimSuffix("123oxo", "xo"))
fmt.Println(strings.TrimPrefix("oxo123", "ox")) /// o123
*/
