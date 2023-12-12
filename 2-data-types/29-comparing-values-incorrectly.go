package datatypes

/*
Mistake 29: Comparing values incorrectly

Devese ficar atento a types primitimos que não são comparaveis, e para eles fazer uso de reflection.

Isso abaixo não funcionará apenas por causa do float
var cust1 any = customer{id: "x", operations: []float64{1.}}
var cust2 any = customer{id: "x", operations: []float64{1.}}
fmt.Println(cust1 == cust2)

A correção seria:
fmt.Println(reflect.DeepEqual(cust1, cust2))

*/
