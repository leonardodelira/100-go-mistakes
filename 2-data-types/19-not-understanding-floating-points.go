package datatypes

/*
Mistake 19: Not understanding floating points
float32 e float64 são aproximações, por isso é importante manter em mante:
- quando comparar dois floating-point, verifique se a diferença está num range aceitavel
- quando for trabalhar com adições e subtração, agrupe as operações em uma ordem de magnitude similar para favorecer a precisão
- se as operações envolver também multiplicações e divisoes, executar elas primeiros para melhor precisão.
*/
