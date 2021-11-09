package models

type DepthFirst struct {
	Grapho   *Grapho
	Solucion []*Solucion
}

func NewDepthFirst(g *Grapho) *DepthFirst {
	dp := &DepthFirst{
		Grapho: g,
	}
	return dp
}

type Solucion struct {
	RecorridoPosMapa []*PosicionMapa
}

func (df *DepthFirst) Resolver() []*Solucion {
	grapho := df.Grapho
	var solucion []Node
	solucion = df.resolver(grapho.root, solucion)
	return df.Solucion
}

func (df *DepthFirst) resolver(node *Node, s []Node) []Node {
	if node != nil {
		//print("->[" + node.Posicion.ValueCenter + "](" + fmt.Sprint(node.Coordenadas.X) + "," + fmt.Sprint(node.Coordenadas.Y) + ")")
		s = append(s, *node)

		valor := node.Posicion.ValueCenter
		if valor == IconObjetivo {
			sol := node.Recorrido
			solucion := &Solucion{RecorridoPosMapa: sol}

			df.Solucion = append(df.Solucion, solucion)
			return s
		} else {

			/*
				t1 := 0.0
				t2 := 0.0
				t3 := 0.0
				t4 := 0.0




					sTop := df.resolver(node.TopNode, s)
					 if sTop == nil {
						t1 = tiempoSolucion(s)
					}

					sLeft := df.resolver(node.LeftNode, s)
					if sTop == nil {
						t2 = tiempoSolucion(s)
					}

					sDown := df.resolver(node.DownNode, s)
					if sDown == nil {
						t3 = tiempoSolucion(s)
					}

					sRigth := df.resolver(node.RigthNode, s)
					if sRigth == nil {
						t4 = tiempoSolucion(s)
					}

					max := math.Max(math.Max(t1, t2), math.Max(t1, t2))
					if max == t1 {
						return sTop
					} else if max == t2 {
						return sLeft
					} else if max == t3 {
						return sDown
					} else if max == t4 {
						return sRigth
					}

			*/

			df.resolver(node.TopNode, s)
			df.resolver(node.LeftNode, s)
			df.resolver(node.DownNode, s)
			df.resolver(node.RigthNode, s)

			//left := resolver(node.LeftNode, i+1, s)
			//down := resolver(node.DownNode, i+1, s)
			//rigth := resolver(node.RigthNode, i+1, s)

		}
	}
	return nil
}

/*
func (g *Grapho) toNodeArray(rec []*PosicionMapa) []*Node {
	var recNode []*Node
	node := g.root
	recNode = append(recNode, node)
	for i := 0; i < len(rec); i++ {

		if node.TopNode != nil {
			if node.TopNode.Coordenadas.Equals(&rec[i].Coordenadas) {
				node = node.TopNode
			}
		} else if node.DownNode != nil {
			if node.DownNode.Coordenadas.Equals(&rec[i].Coordenadas) {
				node = node.DownNode
			}
		} else if node.LeftNode != nil {
			if node.LeftNode.Coordenadas.Equals(&rec[i].Coordenadas) {
				node = node.LeftNode
			}
		} else if node.RigthNode != nil {
			if node.RigthNode.Coordenadas.Equals(&rec[i].Coordenadas) {
				node = node.RigthNode
			}
		}

		recNode = append(recNode, node)
	}

	return recNode
}

*/
func TiempoSolucion(s []*PosicionMapa) float64 {
	tiempo := 0.0
	totalTiempo := 0.0
	OrientacionActual := 0
	for i := 0; i < len(s); i++ {
		terreno := s[i].ValueCenter
		tiempoTerreno := obtenerTiempoTraslado(terreno)
		tiempoGiro := 0.0
		if i != 0 {
			xi, yi := s[i-1].Coordenadas.GetCoordenadas()
			xf, yf := s[i].Coordenadas.GetCoordenadas()
			giro := 0
			//si se mantiene fijo en x
			if xi == xf {
				//si avanza a la derecha
				if yi < yf {
					giro = OrientacionRigth
				} else {
					giro = OrientacionLeft
				}
				//si se mentiene fijo en y
			} else {
				//si se mueve hacia abajo
				if xi < xf {
					giro = OrientacionDown
				} else {
					giro = OrientacionTop
				}

			}
			tiempoGiro = obtenerTiempoGiro(OrientacionActual, giro)
			OrientacionActual = giro
		}

		tiempo += tiempoTerreno + tiempoGiro
		totalTiempo = totalTiempo + tiempo
	}

	return totalTiempo
}

/*

func TiempoSolucion(s []*Node) float64 {
	tiempo := 0.0
	totalTiempo := 0.0
	for i := 0; i < len(s); i++ {
		terreno := s[i].Posicion.ValueCenter
		tiempoTerreno := obtenerTiempoTraslado(terreno)
		tiempoGiro := 0.0
		if i != 0 {
			tiempoGiro = obtenerTiempoGiro(s[i-1].Orientacion, s[i].Orientacion)
		}

		tiempo += tiempoTerreno + tiempoGiro
		totalTiempo = totalTiempo + tiempo
	}

	return totalTiempo
}

*/
