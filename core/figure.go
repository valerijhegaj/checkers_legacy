package core

type Figure interface {
	GetOwnerId() int
	// in each Figure implemented rules where you can move
	Move(desk *Field, actualPosition Coordinate, newPosition ...Coordinate) (bool, Coordinate)
	IsMoveOne(desk *Field, actualPosition, newPosition Coordinate) bool
}
