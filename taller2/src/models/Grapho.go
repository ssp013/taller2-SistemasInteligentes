package models

type Grapho struct {
	root *Node
}

type Node struct {
	Coordenadas *Coordenadas
	Posicion    *PosicionMapa
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

	g.root = &Node{
		Coordenadas: spirit.Estado.Coordenadas,
		Posicion:    pos,
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
			if !posTop.ObstacleDown {
				node.TopNode = &Node{
					Coordenadas: NewCoordenadas(xTop, yTop),
					Posicion:    posTop,
				}
			} else {
				node.TopNode = nil
			}
		}
		//OBSTACLES
		//PARA IR HACIA IZQUIERDA
		xLeft := x
		yLeft := y - 1

		if yLeft >= 0 {
			posLeft := sup.Mapa[xLeft][yLeft]
			if posLeft.ObstacleLeft {
				node.LeftNode = &Node{
					Coordenadas: NewCoordenadas(xLeft, yLeft),
					Posicion:    posLeft,
				}
			} else {
				node.LeftNode = nil
			}
		}

		//PARA IR HACIA ABAJO
		xDown := x + 1
		yDown := y

		if xDown < sup.TamanoN {
			posDown := sup.Mapa[xDown][yDown]
			if !posDown.ObstacleTop {
				node.DownNode = &Node{
					Coordenadas: NewCoordenadas(xDown, yDown),
					Posicion:    posDown,
				}
			} else {
				node.DownNode = nil
			}

		}
		//PARA IR HACIA DERECHA
		xRigth := x
		yRigth := y + 1

		if yRigth < sup.TamanoM {
			posRigth := sup.Mapa[xRigth][yRigth]
			if !posRigth.ObstacleLeft {
				node.RigthNode = &Node{
					Coordenadas: NewCoordenadas(xRigth, yRigth),
					Posicion:    posRigth,
				}
			} else {
				node.RigthNode = nil
			}

		}

		generateNodes(node.TopNode, sup)
		generateNodes(node.LeftNode, sup)
		generateNodes(node.DownNode, sup)
		generateNodes(node.RigthNode, sup)
	}

}

/*
func notIn(source int, pasos []int) bool {
	estado := false
	for i := 0; i < len(pasos); i++ {
		if source == pasos[i] {
			estado = true
		}
	}
	return estado
}

func recursive_dfs(g *Grapho, source int, l []int) []int {
	estado := notIn(source, l)
	if !estado {
		l = append(l, source)

	}
	return l
}
*/
