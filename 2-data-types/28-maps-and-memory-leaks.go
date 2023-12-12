package datatypes

/*
Mistake 28: Maps and memory leaks

Adicionar valores em um map faz com que esse map cresça, porém quando nós deletamos os valores, os "buckets"
criados para esse map continuarão a existir e isso consumirá memória.

Então por exemplo:
Se temos o seguinte map "map[int][128]byte" e adicionamos 1 milhao de elementos nele, ele irá consumir de memória
aproximadamente 461MB. E depois de remover todos os elementos ele irá estar consumindo 293MB.
Isso ocorre porque depois que o map cresceu, ele nunca voltará a ser como antes, porque os buckets criado para o map se manteram ativos.

Uma maneira de amenizar um pouco essa situação é criar um map onde seu "value" seja um ponteiro, exemplo:

map[int]*[128]byte

Dessa forma, se fazemos o mesmo teste citado acima quando removermos todos os elementos desse map ele estará consumindo apenas
38MB de memória. Isso ocorre porque os buckets estavam armazenando apenas a referencia de valores e não o valor de 120bytes em si.
	Com isso o "tamanho do bucket" era menor.
*/
