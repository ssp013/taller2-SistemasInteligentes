package models

type PosicionMapa struct {
	Coordenadas   Coordenadas
	ValueCenter   string
	ObstacleTop   bool
	ObstacleLeft  bool
	ObstacleDown  bool
	ObstacleRigth bool
}
