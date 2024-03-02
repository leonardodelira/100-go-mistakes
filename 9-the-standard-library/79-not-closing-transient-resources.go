package thestandardlibrary

/*
Mistake 79: Not closing transient resources

Aqui basicamente vimos a importancia de fazer uso do metodo .Close() para as structs que implementam
a interface io.Closer().
Caso a gente não feche os recursos que utilizamos sempre teremos aquela memória alocada sem uso que irá
impactar no performance de nosso app.

Resumindo, devo ficar atento se a struct que estou trabalhando implementa a interface io.Closer(). Se a struct
implementar significa que ela possui o metodo .Close().
Com isso devo ter em mente que eventualmente preciso fechar o uso de recursos.
*/
