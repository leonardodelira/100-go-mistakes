package datatypes

import "fmt"

/*
Mistake 25: Unexpected side effects using slice append

Temos que ficar atentos quando fazemos 'slicing' porque normalmente os slices vão apontar para o mesmo backing array
*/

func ExampleMistake25() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s3 := append(s2, 10)

	//output: s1=[1 2 10], s2=[2], s3=[2 10]
	/*
		O valor de s1 foi alterado por causa do append feito no s3, isso ocorre porque eles estavam apontando para o mesmo
		backing array.
		Explicação: criamos o slice1 e fizemos um slice para o slice2, neste momento o slice2 tem o length de 1 porém o capacity de 2.
		quando é feito o append para s3, o último capacity que sobrou é preenchido. Porém nesse cenário como o slice2 está apontando
		para o mesmo backing array de s1, o último valor de slice1 será alterado.
	*/

	fmt.Print(s3)
}

/*
Umas das possibilidades para resolver isso, é criar uma cópia do slice orignal.
Com o lenght limitado ao valor que precisamos.
*/
func ExampleMistake25_CorrectSolution1() {
	s1 := []int{1, 2, 3}
	var s2 = make([]int, 2)
	copy(s2, s1)
	s3 := append(s2, 10)

	fmt.Println("s1", s1) //[1 2 3]
	fmt.Println("s2", s2) //[1 2]
	fmt.Println("s3", s3) //[1 2 10]
}

/*
Outra solução um pouco mais direta onde não precisamos fazer uso do copy.
É por fazer um slicing limitando o tamanho de nosso capacity.
No exemplo abaixo nós pegamos os dois primeiros elementos do slice1 e limitamos esse novo slice2 com capacity de 2 também.
Porque isso? Quando um novo append for feito, um novo backing array será criado porque o nosso capacity chegou ao limite.
E agora teremos um backing array para cada slice.
*/
func ExampleMistake25_CorrectSolution2() {
	s1 := []int{1, 2, 3}
	s2 := s1[:2:2]
	s2 = append(s2, 10)

	fmt.Println("s1", s1) //[1 2 3]
	fmt.Println("s2", s2) //[1 2, 10]
}
