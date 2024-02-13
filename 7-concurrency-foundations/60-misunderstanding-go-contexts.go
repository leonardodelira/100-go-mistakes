package concurrencyfoundations

/*
Mistake 60: Misunderstanding Go contexts

Um Context carrega um deadline, um sinal de cancelamento e outros valores atraves dos limites da API.

deadline = uma atividade pode ser pausada se o "dealine" chegou ao fim.
	por exemplo um I/O request ou uma goroutine esperando por mensagens de um channel.
	ex: context.WithTimeout(context.Background(), 4*time.Second)

cancellation signals = quando a função cancel() do contexto ser chamada, todos os outros processos que tiverem acesso
a esse contexto podem reagir ao sinal de cancelamento.
	ex: ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

context values = contexto onde podemos carregar uma lista key-value. Podemos usar context value para passar valores a outra camadas
da nossa aplicação de forma mais limpa. Por exemplo, em um api rest podemos ter um middleware onde esse middlware faz alguma validação e envia
o resultado no contexto para o handler.
	ex: ctx := context.WithValue(parentCtx, "key", "value")
*/
