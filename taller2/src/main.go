package main
import ""
import (
	models "taller2/src/models"
)

/* con go.mod podemos hacer el import de arriba */
func main() {

	reglas := &models.Reglas{} /* := Me asigna el tipo automatico */
	reglas.Init()              /* eL INIT se hace solamente para que se le asignaran los valores constantes. */

	superficie := &models.Superficie{}
	superficie.Init(reglas)
	/* n  y m : reciben tamanio n y m */
	n := superficie.TamanoN
	m := superficie.TamanoM
	/* ACá quedamos!!! */
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			superficie.Mapa[i][j] = "."
		}
	}
	s := &models.Spirit{}
	spirit := s.NewSpitit()

	sx := spirit.Estado.Coordenadas.X
	sy := spirit.Estado.Coordenadas.Y
	superficie.Mapa[sx][sy] = spirit.Icon
	superficie.Print()

}
