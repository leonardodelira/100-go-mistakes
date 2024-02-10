package errormanagement

import "fmt"

/*
Mistake 52: Handling an error twice

Nós devemos logar um erro ou retornar um erro, nunca os dois.
Devemos fazer uso do error wrapper que o Go nos fornece para adicionar mais contexto em nossos erros.

Lembre-se, "logar" um erro é uma forma de manusear o erro (handling error). Então devemos logar um erro ou retorna-lo.
Nunca os dois.
*/

type Route struct{}

/*
No exemplo abaixo nós validamos a latitude e longitude na mesma função.
Porém é apenas no nosso "caller" que adicionamos mais contexto ao erro, fazendo uso do %w.
Adicionar esse contexto nos ajuda a dar mais informações sobre onde ocorreu o erro e também
evita o uso de dois logs.
*/
func GetRoute(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates(srcLat, srcLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate latitude coordinates: %w",
				err)
	}

	err = validateCoordinates(dstLat, dstLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate longitude coordinates: %w",
				err)
	}

	return Route{}, nil
}

func validateCoordinates(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}
