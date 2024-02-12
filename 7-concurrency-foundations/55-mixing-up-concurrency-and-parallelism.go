package concurrencyfoundations

/*
Mistake 55: Mixing up concurrency and parallelism

Parallelism: partes do sistema independentes. Talvez executando a mesma tarefa, porém de forma independente.
Concurrency: é sobre lidar com várias coisas ao mesmo tempo. E para isso fazemos uso do paralelismo.

Ilustração da cafeteria:

Paralelismo:
Digamos que temos uma cafeteria, existe uma fila de pedidos. A pessoa faz seu pedido e espera ficar pronto para receber sem sair da fila.
	Ou seja, o atendente recebe o pedido e também o prepara.
Para ganhar velocidade no atendimento e preparo do café, outra fila é criada com um novo atendente e uma nova máquina de café.
Agora temos dois processos independentes sendo executados ao mesmo tempo. Isso é paralelismo.

Concorrencia:
Para ganhar agilidade, quando a pessoa faz seu pedido ela vai para outra fila esperar seu pedido ficar pronto. Dessa forma a fila de pessoas
que ainda vão fazer o pedido começa a andar mais rapido.
	E agora temos dois funcionarios, um atendente e outro que prepara o café.
	Isso é concorrencia, a estrutura mudou.
Com essa nova abordagem podemos evoluir o processo de forma independentes, como assim?
Digamos que o processo de preparação de café precisa ser mais rapido, basta colocarmos outro funcionário para fazer os cafés.
	Nesse ponto temos dois funcionários fazendo os cafés, temos um paralelimos. (dois processos independentes).
	Porém por meio da concorrencia nós "manuseamos" esses processos pararelos para ter sincronia.

	"Concorrencia significa lidar com muitas coisas ao mesmo tempo. Paralelismo é fazer muitas coisas ao mesmo tempo"
*/
