package models

type Spirit struct {
	Icon   string
	Estado *Estado
}

//gira hacia arriba y retorna el tiempo que demoro

func NewSpitit() *Spirit {

	estado := &Estado{
		Coordenadas: NewCoordenadasRandom(N, M),
		Orientacion: OrientacionRigth,
	}

	spirit := &Spirit{
		Icon:   "▶",
		Estado: estado,
	}

	return spirit
}

/*

func (s *Spirit) GirarDown() {
	s.Icon = "▼"
	s.Estado.Orientacion = "down"
}

func (s *Spirit) GirarUp() {
	s.Icon = "▲"
	s.Estado.Orientacion = "up"
}

func (s *Spirit) GirarRigth() {
	s.Icon = "▶"
	s.Estado.Orientacion = "rigth"
}

func (s *Spirit) GirarLeft() {
	s.Icon = "◀"
	s.Estado.Orientacion = "left"
}

*/
