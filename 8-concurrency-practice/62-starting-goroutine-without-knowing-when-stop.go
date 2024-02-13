package concurrencypractice

/*
Mistake 62: Starting a goroutine without knowing when to stop it

Inicar uma goroutine sem saber quando ela irá ser finalizar, é um problema de desing.
Uma goroutine é um recurso como qualquer outro e precisar ser finalizado quando necessário para liberar memória ou outros recursos.

Abaixo um exemplo de como podemos finalizar uma goroutine de maneira "greatfull".
Quando o processo do main finalizar, chamamos o metodo de close da nossa struct.
*/

/* Some resources */
type watcher struct{}

func (w watcher) watch() {}

func main() {
	w := newWatcher()
	defer w.close()

	// Run the application
}

func newWatcher() watcher {
	w := watcher{}
	go w.watch()
	return w
}

func (w watcher) close() {
	// Close the resources
}
