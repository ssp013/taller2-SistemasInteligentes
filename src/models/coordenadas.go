package models

type Coordenadas struct {
	X int
	Y int
}

func (*Coordenadas) NewCoordenadas(tam_n int, tam_m int) *Coordenadas {
	c := &Coordenadas{
		X: 5,
		Y: 3,
	}
	return c
}
