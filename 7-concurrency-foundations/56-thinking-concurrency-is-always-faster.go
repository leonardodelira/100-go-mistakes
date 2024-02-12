package concurrencyfoundations

/*
Mistake 56: Thinking concurrency is always faster

Devemos ter em mente que trabalhar que concorrencia ou/e paralelismo nem sempre é mais vantajoso.
Criar novas goroutines é ligeiramente custoso, e também tem a questão do swtich context.
Então devemos ser cuidadosos ao querer criar novas gorountines para um fluxo de trabalho pequeno. Muitas vezes o custo
de criar e manusear essas novas goroutines é maior do que se mantivermos o processamento em apenas uma goroutine.
*/
