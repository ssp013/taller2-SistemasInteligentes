package models

/* Constantes globales */
const (
	N        = 3
	M        = 3
	Giro     = 4.0
	VelAbr   = 0.5
	VelLlano = 1.2
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
