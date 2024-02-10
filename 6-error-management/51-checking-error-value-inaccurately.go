package errormanagement

/*
Mistake 51: Checking an error value inaccurately

"Sentinel error" é um erro global. Normalmente fazemos uso desse tipo de erro para errors previsíveis.
Como por exeplo:

sql.ErrNoRows - nenhum valor foi encontrado na query
io.EOF - não tem mais input para ler

Fazemos uso de sentinel errors em nosso código para tratar de maneira correta os errors que recebemos.
Aqui o ponto de atenção que temos é, da mesma forma que fazemos uso de errors.As() para ver o tipo de erro que recebemos
de forma recursiva. Podemos fazer uso de errors.Is para ver o "value" do error que recebemos, porque ele pode vir
dentro de um wrapper também.

err := query()
if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
        // ...
    } else {
        // ...
    }
}
*/
