package core

type Figure interface {
	GetOwnerId() int
	// in each Figure implemented rules where you can move
	Move(desk *Field, from Coordinate, way []Coordinate) (bool, Coordinate)
	IsMoveOne(desk *Field, from, to Coordinate) (bool, Coordinate)
}
