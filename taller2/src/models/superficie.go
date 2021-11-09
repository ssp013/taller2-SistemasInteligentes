package models

import (
	"fmt"
)

/* mÉTODOS INICIO CON MAYUSCULA SON PUBLICOS , CON MINUSCULA SON PRIVADOS */
const (
	//objetos
	IconLlano     = "."
	IconAbr       = "∷"
	IconObstaculo = "│"
	IconObjetivo  = "Ⓧ"
)

type Superficie struct {
	Mapa    [N][M]*PosicionMapa
	TamanoN int
	TamanoM int
	Reglas  *Reglas
}

func (s *Superficie) Init(r *Reglas) {
	var mapa [N][M]*PosicionMapa /* crea una matriz de tipo string */
	/* N: Fila M: es columna */
	s.TamanoN = N
	s.TamanoM = M
	s.Mapa = mapa
	s.Reglas = r

	s.crearPosiciones()
	s.posicionarTerrenoAbrupto(N)
	s.posicionarObjetivo()
	s.posicionarObstaculos(N)
}

func (s *Superficie) crearPosiciones() {
	for i := 0; i < s.TamanoN; i++ {
		for j := 0; j < s.TamanoM; j++ {
			p := &PosicionMapa{ValueCenter: IconLlano, Coordenadas: *NewCoordenadas(i, j)}
			s.Mapa[i][j] = p
		}
	}
}

func (s *Superficie) posicionarTerrenoAbrupto(cantidad int) {
	cantidadAbrupto := randNum(cantidad) + cantidad/2
	for i := 0; i < cantidadAbrupto; i++ {
		c := NewCoordenadasRandom(N, M)
		x, y := c.GetCoordenadas()
		for s.Mapa[x][y].ValueCenter == IconAbr {
			c = NewCoordenadasRandom(N, M)
			x, y = c.GetCoordenadas()
		}
		s.Mapa[x][y].ValueCenter = IconAbr
	}
}

func (s *Superficie) posicionarObjetivo() {
	c := NewCoordenadasRandom(N, M)
	x, y := c.GetCoordenadas()

	for s.Mapa[x][y].ValueCenter != IconLlano {
		c = NewCoordenadasRandom(N, M)
		x, y = c.GetCoordenadas()
	}
	s.Mapa[x][y].ValueCenter = IconObjetivo
}

func (s *Superficie) PosicionarSpirit(spirit *Spirit) {
	c := NewCoordenadasRandom(N, M)
	x, y := c.GetCoordenadas()

	for s.Mapa[x][y].ValueCenter != IconLlano {
		c = NewCoordenadasRandom(N, M)
		x, y = c.GetCoordenadas()
	}
	s.Mapa[x][y].ValueCenter = spirit.Icon
	spirit.Estado.Coordenadas = c

}
func (s *Superficie) posicionarObstaculos(cantidad int) {
	cantidadVertical := randNum(cantidad) + cantidad/2

	for i := 0; i < cantidadVertical; i++ {
		c := NewCoordenadasRandom(N, M)
		x, y := c.GetCoordenadas()

		for s.Mapa[x][y].ObstacleRigth {
			c = NewCoordenadasRandom(N, M)
			x, y = c.GetCoordenadas()
		}
		//crea un obstaculo derecho
		s.Mapa[x][y].ObstacleRigth = true
		yRigth := y + 1
		if y+1 < s.TamanoM {
			s.Mapa[x][yRigth].ObstacleLeft = true
		}

	}

	cantidadHorizontal := randNum(cantidad) + cantidad/2
	for j := 0; j < cantidadHorizontal; j++ {
		c := NewCoordenadasRandom(N, M)
		x, y := c.GetCoordenadas()

		for s.Mapa[x][y].ObstacleDown {
			c = NewCoordenadasRandom(N, M)
			x, y = c.GetCoordenadas()
		}
		//crea un obstaculo inferior
		s.Mapa[x][y].ObstacleDown = true
		xDown := x + 1
		if x+1 < s.TamanoN {
			s.Mapa[xDown][y].ObstacleTop = true
		}

	}
}

/* PRINTEO */
func (s *Superficie) Print() {

	for a := 0; a < s.TamanoM; a++ {
		fmt.Print("\t", a)
	}
	fmt.Println()
	for i := 0; i < s.TamanoN; i++ {
		fmt.Print(" ", i, "\t")
		for j := 0; j < s.TamanoM; j++ {
			value := fmt.Sprint(s.Mapa[i][j].ValueCenter)
			if s.Mapa[i][j].ObstacleRigth {
				fmt.Print(value, "   ", IconObstaculo+"\t")
			} else {
				fmt.Print(value, "\t")
			}
		}
		fmt.Print("\n\t")
		for k := 0; k < s.TamanoM; k++ {
			if s.Mapa[i][k].ObstacleDown {
				fmt.Print("ㅡㅡ", "\t")
			} else {
				fmt.Print("", "\t")
			}

		}
		fmt.Println("")
		fmt.Println("")

	}
}
