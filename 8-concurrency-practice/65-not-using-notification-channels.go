package concurrencypractice

/*
Mistake 65: Not using notification channels

Utilizamos channels para se comunicar entre as goroutines.
E um channel pode existir com ou sem dados. Por exemplo, podemos criar um channel apenas para notificar que algo foi disconectado.
ex: disconectCh make(chan ...)

Nesse casos onde o channel não precisa transportar nenhum tipo de dado, o ideial é criar um channel do tipo struct{}, porque uma
strutct vazia não ocupa bytes em memória. (diferente da interface por exemplo)

disconectCh := make(chan struct{}).

Resumindo: se formos criar um channel apenas para notificar alguma ação onde nenhum tipo de dado é necessário ser enviado. Criamos
esse channel com o tipo struct{}
*/
