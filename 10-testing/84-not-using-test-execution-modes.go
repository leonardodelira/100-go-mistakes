package testing

/*
Mistake 84: Not using test execution modes

parallel:
	Podemos executar tests em pararelos, isso pode ser útil para testes de longa duração. Basta marcar nosso teste com t.Parallel().
	E na hora de executar os testes podemos passar a flag -parallel + quantidade de testes rodando em paralelo:
	ex "-parallel 16"
shuffle:
	Essa flag executa nossos testes de maneira aleatória, isso pode ser útil para detectar dependencias escondidas em nosso código, por exemplo quando um teste depende de outro para funcionar
	quando na verdade deveriam funcionar de modo isolado.
*/

//cmd: go test -shuffle=on -v .

/*
A flag -v nos permite saber o "id" do testes aleatórios que foram executados. Isso é útil caso a gente encontre algum
erro durante a execução dos testes e tentamos reproduzir novamente.
Com o "id" basta executar os testes novamente passando esse valor como parametro
*/

//cmd go test -shuffle=id_aqui -v .

//Agora a mesma ordem dos testes será executada.
