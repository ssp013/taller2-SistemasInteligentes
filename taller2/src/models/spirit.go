package models

type Spirit struct {
	Icon   string
	Estado *Estado
}

//gira hacia arriba y retorna el tiempo que demoro

func (*Spirit) NewSpitit() *Spirit {

	coordenadas := &Coordenadas{}
	estado := &Estado{
		Coordenadas: coordenadas.NewCoordenadas(N, M),
		Orientacion: "Rigth",
	}

	spirit := &Spirit{
		Icon:   "▶",
		Estado: estado,
	}

	return spirit
}

func (s *Spirit) girarDown() {
	s.Icon = "▼"
}

func (s *Spirit) girarUp() {
	s.Icon = "▲"
}

func (s *Spirit) girarRigth() {
	s.Icon = "▶"
}

func (s *Spirit) girarLeft() {
	s.Icon = "◀"
}
