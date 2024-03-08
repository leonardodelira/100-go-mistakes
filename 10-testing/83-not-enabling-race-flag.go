package testing

/*
Mistake 83: Not enabling the -race flag

A flag "-race" nos ajuda a identificar possíveis data race em nosso código. Podemos usar essa flag ao executar
testes e também ao criar nossa build.
Porém devemos ficar atentos, fazer uso da flag "-race" faz com que aumente o consumo de memória da nossa aplicação
e também ocasiona em perda de performance. Então o ideal é fazer uso dessa flag apenas em ambiente local ou durante
o CI.
*/
