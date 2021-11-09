package models

type Grapho struct {
	root *Node
}

type Node struct {
	Coordenadas *Coordenadas
	Posicion    *PosicionMapa
	Orientacion int
	Recorrido   []*PosicionMapa
	TopNode     *Node
	LeftNode    *Node
	DownNode    *Node
	RigthNode   *Node
}

func NewGrapho() *Grapho {
	return &Grapho{}
}

func (g *Grapho) GenerateGrapho(sup *Superficie, spirit *Spirit) {
	x, y := spirit.Estado.Coordenadas.GetCoordenadas()
	pos := sup.Mapa[x][y]

	var array []*PosicionMapa
	g.root = &Node{
		Coordenadas: spirit.Estado.Coordenadas,
		Posicion:    pos,
		Orientacion: spirit.Estado.Orientacion,
		Recorrido:   append(array, pos),
	}
	generateNodes(g.root, sup)
}

func generateNodes(node *Node, sup *Superficie) {

	if node != nil {
		/* Posision mapa OBJETO, tenemos aca las coordenadas, value, top, left right */
		pos := node.Posicion
		x, y := pos.Coordenadas.GetCoordenadas()
		/*
			X , Y
			TOP 	: X - 1, Y
			LEFT 	: X, Y-1
			DOWN 	: X + 1, Y
			RIGTH 	: X, Y + 1
		*/
		//PARA IR HACIA ARRIBA
		xTop := x - 1
		yTop := y

		if xTop >= 0 {
			posTop := sup.Mapa[xTop][yTop]
			esta := esta(posTop, node.Recorrido)
			if !posTop.ObstacleDown && !esta {
				array := append(node.Recorrido, posTop)
				node.TopNode = &Node{
					Coordenadas: NewCoordenadas(xTop, yTop),
					Posicion:    posTop,
					Orientacion: OrientacionTop,
					Recorrido:   array,
				}
			} else {
				node.TopNode = nil
			}
		} else {
			node.TopNode = nil
		}
		//OBSTACLES
		//PARA IR HACIA IZQUIERDA
		xLeft := x
		yLeft := y - 1

		if yLeft >= 0 {
			posLeft := sup.Mapa[xLeft][yLeft]
			esta := esta(posLeft, node.Recorrido)
			if posLeft.ObstacleRigth && !esta {
				array := append(node.Recorrido, posLeft)
				node.LeftNode = &Node{
					Coordenadas: NewCoordenadas(xLeft, yLeft),
					Posicion:    posLeft,
					Orientacion: OrientacionLeft,
					Recorrido:   array,
				}
			} else {
				node.LeftNode = nil
			}
		} else {
			node.LeftNode = nil
		}

		//PARA IR HACIA ABAJO
		xDown := x + 1
		yDown := y

		if xDown < sup.TamanoN {
			posDown := sup.Mapa[xDown][yDown]
			esta := esta(posDown, node.Recorrido)
			if !posDown.ObstacleTop && !esta {
				array := append(node.Recorrido, posDown)
				node.DownNode = &Node{
					Coordenadas: NewCoordenadas(xDown, yDown),
					Posicion:    posDown,
					Orientacion: OrientacionDown,
					Recorrido:   array,
				}
			} else {
				node.DownNode = nil
			}

		} else {
			node.DownNode = nil
		}
		//PARA IR HACIA DERECHA
		xRigth := x
		yRigth := y + 1

		if yRigth < sup.TamanoM {
			posRigth := sup.Mapa[xRigth][yRigth]
			esta := esta(posRigth, node.Recorrido)
			if !posRigth.ObstacleLeft && !esta {
				array := append(node.Recorrido, posRigth)
				node.RigthNode = &Node{
					Coordenadas: NewCoordenadas(xRigth, yRigth),
					Posicion:    posRigth,
					Orientacion: OrientacionRigth,
					Recorrido:   array,
				}
			} else {
				node.RigthNode = nil
			}

		} else {
			node.RigthNode = nil
		}

		generateNodes(node.TopNode, sup)
		generateNodes(node.LeftNode, sup)
		generateNodes(node.DownNode, sup)
		generateNodes(node.RigthNode, sup)
	}
}

func esta(key *PosicionMapa, list []*PosicionMapa) bool {

	for i := 0; i < len(list); i++ {
		if key != nil && key.Coordenadas.X == list[i].Coordenadas.X && key.Coordenadas.Y == list[i].Coordenadas.Y {
			return true
		}
	}
	return false
}

/*
func recursive_dfs(g *Grapho, source int, l []int) []int {
	estado := notIn(source, l)
	if !estado {
		l = append(l, source)

	}
	return l
}
*/
