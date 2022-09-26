package test

import "chekers/core"

type Test_figure struct {
	OwnerId int
}

func (c Test_figure) GetOwnerId() int {
	return c.OwnerId
}

func (c Test_figure) Move(desk *core.Field, from core.Coordinate, way []core.Coordinate) (bool, core.Coordinate) {
	return false, from
}

func (c Test_figure) IsMoveOne(desk *core.Field, from, to core.Coordinate) (bool, core.Coordinate) {
	return false, desk.BordersLeft
}
