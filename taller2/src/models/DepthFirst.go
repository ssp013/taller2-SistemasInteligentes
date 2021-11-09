package models

import "math"

type DepthFirst struct {
	Grapho *Grapho
}

func NewDepthFirst(g *Grapho) *DepthFirst {
	dp := &DepthFirst{
		Grapho: g,
	}
	return dp
}

func (df *DepthFirst) Resolver() []Node {
	grapho := df.Grapho
	var solucion []Node
	solucion = resolver(grapho.root, solucion)
	return solucion
}

func resolver(node *Node, s []Node) []Node {
	if node != nil {
		s = append(s, *node)
		valor := node.Posicion.ValueCenter
		if valor == IconObjetivo {
			return s
		} else {
			t1 := 0.0
			t2 := 0.0
			t3 := 0.0
			t4 := 0.0

			sTop := resolver(node.TopNode, s)
			/* */ if sTop == nil {
				t1 = tiempoSolucion(s)
			}

			sLeft := resolver(node.LeftNode, s)
			if sTop == nil {
				t2 = tiempoSolucion(s)
			}

			sDown := resolver(node.DownNode, s)
			if sDown == nil {
				t3 = tiempoSolucion(s)
			}

			sRigth := resolver(node.RigthNode, s)
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
			//left := resolver(node.LeftNode, i+1, s)
			//down := resolver(node.DownNode, i+1, s)
			//rigth := resolver(node.RigthNode, i+1, s)

		}
	}
	return nil
}

func tiempoSolucion(s []Node) float64 {
	tiempo := 0.0
	for i := 0; i < len(s); i++ {
		terreno := s[i].Posicion.ValueCenter
		tiempoTerreno := obtenerTiempoTraslado(terreno)
		tiempoGiro := obtenerTiempoGiro(s[i].Orientacion, s[i].Orientacion)
		tiempo += tiempoTerreno + tiempoGiro
	}

	return tiempo
}
