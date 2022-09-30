package core

// in each Figure implemented rules where you can move
type Figure interface {
	GetOwnerId() int
	Move(desk *Field, from Coordinate, way []Coordinate) (bool, Coordinate)
	IsMoveOne(desk *Field, from, to Coordinate) (bool, Coordinate)
	GetAvailableMoves(desk *Field, from Coordinate) []Coordinate
	GetAvailableMovesToEat(desk *Field, from Coordinate) []Coordinate
}
