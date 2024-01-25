package controlstructs

import "os"

/*
Mistake 35: Using defer inside a loop

Devemos tomar cuidado ao user "defer" dentro de um loop.
A função defer é executada quando nossa função principal retorna, ou seja, se colocamos o defer dentro de um loop o defer demorara para ser executado.
Exemplo abaixo.
*/

func BadPraticeDefer(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		//Por exemplo neste caso, o defer só será executado quando o channel ser fechado a chegar o fim da função principal
		//Então teremos vários arquivos abertos e isso pode causar leak memory.
		defer file.Close()
		//Do somethig with the file.
	}
	return nil
}

/*
Neste caso abaixo, criamos uma outra função para lidar com a leitura do arquivo e com o defer.
Agora temos a garantia que o defer será executado em cada iteração do for, porque chamamos outra função e o defer será executado
quando essa outra função ser finalizada.
*/
func GoodPraticeDefer(ch <-chan string) error {
	for path := range ch {
		if err := readPath(path); err != nil {
			return err
		}
	}
	return nil
}

func readPath(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	//do something with file
	return nil
}
