package codeandprojectorganization

/*
Mistake #7 - Returning interfaces

"Accept interfaces, return structs"

A ideia aqui é sempre retornarmos o tipo concreto de nossa implementação e deixar que o client se preocupe em criar as interfaces conforme
necessário. Como assim?

Suponhamos que tenhamos a seguinte estrutura.

//db.go
package db
type Store struct {
   db *sql.DB
}
func NewDB() *Store { ... } //func to initialise DB
func (s *Store) Insert(item interface{}) error { ... } //insert item
func (s *Store) Get(id int) error { ... } //get item by id


//user.go
package user
type UserStore interface {
   Insert(item interface{}) error
   Get(id int) error
}
type UserService struct {
   store UserStore
}
// Accepting interface here!
func NewUserService(s UserStore) *UserService {
   return &UserService{
      store: s,
   }
}
func (u *UserService) CreateUser() { ... }
func (u *UserService) RetrieveUser(id int) User { ... }

Como podemos ver, a struct "UserService" depende de alguns metodos que nós sabemos que existe em "Store". Neste caso, ao invés de "NewDB" ter na sua assinatura
o retorno de uma interface, nós retornamos o valor concreto. Porque? A implementação de interfaces no Go é implicita, então quando injetarmos esse valor
no "NewUserService" irá funcionar sem problema algum visto que ele espera como dependencia qualquer objeto que respeite a assinatura esperada em "UserStore"

Vantagens:
- Desaclopamento
- Flexibilidade
- Fácil de testar

Nossa struct UserService tem acesso apenas aos metodos que necesita, mesmo se no futuro Store tiver mais metodos.
	- E também se novos metodos forem inseridos em Store, nosso código não sofre alteração .. visto que nossa struct UserService necesita ter acesso a apenas o que definirmos na interface UserService.
O teste também é facilitado, visto que vamos precisar mockar apenas os metodos que a struct UserService faz uso.
*/
