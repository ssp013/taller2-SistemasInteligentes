package models

const (
	N        = 20
	M        = 20
	Giro     = 4.0
	VelAbr   = 0.5
	VelLlano = 1.2
)

type Reglas struct {
	TiempoGiro        float32
	VelTerrenoAbrupto float32
	VelTerrenoLano    float32
}

func (r *Reglas) Init() {

	r.TiempoGiro = Giro
	r.VelTerrenoAbrupto = VelAbr
	r.VelTerrenoLano = VelLlano
}
