package models

import "fmt"

type Superficie struct {
	Mapa    [N][M]string
	TamanoN int
	TamanoM int
	Reglas  *Reglas
}

func (s *Superficie) Init(r *Reglas) {
	var mapa [N][M]string
	s.TamanoN = N
	s.TamanoM = M
	s.Mapa = mapa
	s.Reglas = r

}

func (s *Superficie) Print() {
	for i := 0; i < s.TamanoN; i++ {
		for j := 0; j < s.TamanoM; j++ {
			fmt.Print(s.Mapa[i][j], "\t")
		}
		fmt.Println("")
	}
}
