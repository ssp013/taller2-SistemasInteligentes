package models

import (
	"time"
)

type Coordenadas struct {
	X int
	Y int
}

func NewCoordenadasRandom(tam_n int, tam_m int) *Coordenadas {
	c := &Coordenadas{
		X: randNum(tam_n),
		Y: randNum(tam_m),
	}
	return c
}

func NewCoordenadas(x int, y int) *Coordenadas {
	c := &Coordenadas{
		X: x,
		Y: y,
	}
	return c
}

func (c *Coordenadas) GetCoordenadas() (int, int) {
	return c.X, c.Y
}

func (c *Coordenadas) Equals(coordenadas *Coordenadas) bool {
	if c.X == coordenadas.X && c.Y == coordenadas.Y {
		return true
	}
	return false
}

func randNum(limit int) int {
	n := time.Now().Nanosecond() / (time.Now().Second() + 1)
	return n % limit
}
