package main

import (
	"fmt"
	models "taller2/src/models"
	"time"
)

/* con go.mod podemos hacer el import de arriba */
func main() {

	inicio := time.Now()
	reglas := &models.Reglas{} /* := Me asigna el tipo automatico */
	reglas.Init()              /* eL INIT se hace solamente para que se le asignaran los valores constantes. */

	superficie := &models.Superficie{}
	superficie.Init(reglas)

	/* n  y m : reciben tamanio n y m */
	//n := superficie.TamanoN
	//m := superficie.TamanoM
	/* ACá quedamos!!! */

	/*
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				superficie.Mapa[i][j] = "."
			}
		}
	*/
	spirit := models.NewSpitit()
	//spirit.GirarRigth()
	superficie.PosicionarSpirit(spirit)

	superficie.Print()

	grapho := models.NewGrapho()
	grapho.GenerateGrapho(superficie, spirit)

	depthFirst := models.NewDepthFirst(grapho)
	depthFirst.Resolver()
	solucion := depthFirst.Solucion

	/*
		for i := 0; i < len(solucion); i++ {
			rec := solucion[i].RecorridoPosMapa
			print(fmt.Sprint(i))
			tiempo := models.TiempoSolucion(rec)
			for j := 0; j < len(rec); j++ {

				x, y := rec[j].Coordenadas.GetCoordenadas()
				print("->[" + rec[j].ValueCenter + "](" + fmt.Sprint(x) + "," + fmt.Sprint(y) + ")")
			}
			println("\tt:" + fmt.Sprint(tiempo))
		}*/

	indexSolucion := 0
	mejorTiempo := 100000.0

	for k := 0; k < len(solucion); k++ {
		rec := solucion[k].RecorridoPosMapa
		tiempo := models.TiempoSolucion(rec)
		if tiempo > 0 && (tiempo < mejorTiempo) {
			indexSolucion = k
			mejorTiempo = tiempo
		}
	}

	mejorTiempo = mejorTiempo + (1 / models.VelLlano)

	println("\nLA MEJOR SOLUCION ES LA SIGUIENTE: ")

	mejorSolucion := solucion[indexSolucion].RecorridoPosMapa
	for a := 0; a < len(mejorSolucion); a++ {
		x, y := mejorSolucion[a].Coordenadas.GetCoordenadas()
		print("->[" + mejorSolucion[a].ValueCenter + "](" + fmt.Sprint(x) + "," + fmt.Sprint(y) + ")")
	}

	println("CON UN TIEMPO DE: " + fmt.Sprint(mejorTiempo) + " SEGUNDOS")
	println()
	println()

	ejecucion := time.Since(inicio).Milliseconds()
	println("Tiempo de ejecucion " + fmt.Sprint(superficie.TamanoN) + "x" + fmt.Sprint(superficie.TamanoM) + ": " + fmt.Sprint(ejecucion) + "ms")

	/*
		for i := 0; i < len(solucion); i++ {
				rec := solucion[i].Recorrido
				for j := 0; j < len(rec); j++ {
					x, y := rec[j].Coordenadas.GetCoordenadas()
					print("->[" + rec[j].Posicion.ValueCenter + "](" + fmt.Sprint(x) + "," + fmt.Sprint(y) + ")")
				}
				println()
			}
	*/

}
