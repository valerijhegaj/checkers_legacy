package test

import "chekers/core"

type TestFigure struct {
	OwnerId int
}

func (c TestFigure) GetOwnerId() int {
	return c.OwnerId
}

func (c TestFigure) Move(desk *core.Field, from core.Coordinate, way []core.Coordinate) (bool, core.Coordinate) {
	return false, from
}

func (c TestFigure) IsMoveOne(desk *core.Field, from, to core.Coordinate) (bool, core.Coordinate) {
	return false, desk.BordersLeft
}

func (c TestFigure) GetAvailableMoves(desk *core.Field, from core.Coordinate) []core.Coordinate {
	return nil
}

func (c TestFigure) GetAvailableMovesToEat(ddesk *core.Field, from core.Coordinate) []core.Coordinate {
	return nil
}
