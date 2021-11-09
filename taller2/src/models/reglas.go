package models

import "math"

/* Constantes globales */
const (
	N        = 7
	M        = 7
	Giro     = 4
	VelAbr   = 0.5
	VelLlano = 1.2
	//sentido spirit
	OrientacionTop   = 90
	OrientacionLeft  = 180
	OrientacionDown  = 270
	OrientacionRigth = 0
)

/* Struct es decir q atributos tendra esa estructura */
type Reglas struct {
	TiempoGiro        float32
	VelTerrenoAbrupto float32
	VelTerrenoLano    float32
}

/* Constructor, no es necesario que lleve, es por temas de automatitation */
func (r *Reglas) Init() {

	r.TiempoGiro = Giro
	r.VelTerrenoAbrupto = VelAbr
	r.VelTerrenoLano = VelLlano
}

func obtenerTiempoGiro(i int, f int) float64 {
	grado := math.Abs(float64(i - f))
	if grado == 270 {
		grado = 90
	}
	resultado := (grado / 90) * Giro
	return resultado
}

func obtenerTiempoTraslado(terreno string) float64 {
	if terreno == IconLlano {
		return (1 / VelLlano)
	} else if terreno == IconAbr {
		return (1 / VelAbr)
	}
	return 0
}
