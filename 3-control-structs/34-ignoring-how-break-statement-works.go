package controlstructs

import "fmt"

/*
Mistake 34: Ignoring how the break statement works

Devemos ter atenção ao usar "break" dentro de um loop em conjunto com um swtich.
Exemplo abaixo
*/

func LoopSwitchBreakWrong() {
	for i := 0; i < 5; i++ {
		switch i {
		case 2:
			break //Neste caso o break server apenas para o switch e não para o loop.
		default:
			fmt.Println(i)
		}
	}
}

func LoopSwitchBreakCorrect() {
loop:
	for i := 0; i < 5; i++ {
		switch i {
		case 2:
			break loop //A maneira correta de dar break no loop neste caso é por usar um "label", dessa forma fazemos referencia ao loop e nao ao switch
		default:
			fmt.Println(i)
		}
	}
}
