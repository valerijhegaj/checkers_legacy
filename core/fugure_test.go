package core

import "testing"

func testFigure_GetAvailableMoves(
	t *testing.T,
	field *Field,
	from Coordinate,
	availableMoves map[Coordinate]bool,
	messeage string,
	GetAvailableMoves func(desk *Field, from Coordinate) []Coordinate,
) {
	moves := GetAvailableMoves(field, from)
	for _, move := range moves {
		if !availableMoves[move] {
			t.Error(messeage)
		}
		delete(availableMoves, move)
	}
	if len(availableMoves) != 0 {
		t.Error(messeage)
	}
}
