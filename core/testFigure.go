package core

type test_figure struct {
	ownerId int
}

func (c test_figure) GetOwnerId() int {
	return c.ownerId
}

func (c test_figure) Move(desk *Field, actualPosition Coordinate, newPosition []Coordinate) (bool, Coordinate) {
	return false, actualPosition
}

func (c test_figure) IsMoveOne(desk *Field, actualPosition, newPosition Coordinate) bool {
	return false
}
