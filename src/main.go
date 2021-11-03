package main

import (
	models "taller2/src/models"
)

func main() {

	reglas := &models.Reglas{}
	reglas.Init()

	n := reglas.TamanoN
	m := reglas.TamanoM

	superficie := &models.Superficie{}
	superficie.Init(reglas)

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
