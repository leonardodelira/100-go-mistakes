package concurrencypractice

import "sync"

/*
Mistake 74: Copying a sync type

Devemos ficar atentos ao criar cópias de structs que fazem uso do package "sync".
Porque nesse caso mesmo fazendo uso correto do "sync" para proteger o dado de um possível data race, nós teremos
problemas.
Exemplo abaixo:
*/

type Mistake74Account struct {
	mu       sync.Mutex
	balances map[string]int
}

func NewAccountMistake74() Mistake74Account {
	return Mistake74Account{
		balances: map[string]int{},
	}
}

/*
Aqui temos um problema, nosso receiver é do tipo "value" ou seja está sendo criado uma cópia de Account.
Dessa maneira o "mu" não terá acesso ao se outras goroutines estão tentando alterar esse dado.
Por fim, temos erro de data race.
*/
func (a Mistake74Account) Mistake74_CopyingSyncType(key string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balances[key]++
}

/*
Umas das soluções é por alterar o nosso receiver para ser do tipo pointer. Agora o "mu" terá acesso
a outras goroutines que estão tentando acessar esse mesmo dado e irá conseguir evitar erro de "data race".
*/
func (a *Mistake74Account) Mistake74_CorrectSyncType(key string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balances[key]++
}

/*
Outra solução é por alterar o tipo do sync.Mutex em nossa struct, criando ele como um ponteiro para sync.Mutex.
Exemplo abaixo:
*/

// Agora temos um ponteiro para um tipo sync.Mutex, dessa forma nosso receiver pode ser do tipo value sem problemas algum.
type Mistake74Account2 struct {
	mu       *sync.Mutex
	balances map[string]int
}

/*
Visto que agora o "mu" é do tipo ponteiro para sync.Mutex{}, necesitamos iniciar ele na criação da struct.
Caso contrário quando formos fazer uso dele (mu.Lock()) vamos receber um panic porque ele não foi inicializado e então terá o valor
de "nil".
*/
func NewAccount2Mistake74() Mistake74Account2 {
	return Mistake74Account2{
		mu:       &sync.Mutex{},
		balances: map[string]int{},
	}
}

// Agora podemos ter um receiber do tipo value
func (a Mistake74Account2) Mistake74_CopyingSyncType(key string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balances[key]++
}
