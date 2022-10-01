package gamer

import (
	"testing"

	"chekers/core"
)

func TestGamer_GetField(t *testing.T) {
	var c core.GameCore
	field := core.NewField()
	field.BordersRight = core.Coordinate{-1, -1}
	field.BordersLeft = core.Coordinate{-2, -2}
	ptr := core.Coordinate{-1, -1}
	field.Put(ptr, core.TestFigure{0})
	c.InitField(field)

	gamer := Gamer{0, &c}
	fieldInCore := gamer.GetField()
	if field.BordersRight != fieldInCore.BordersRight ||
		len(field.Figures) != 1 ||
		field.At(ptr) != fieldInCore.At(ptr) {
		t.Error("Not that field in core")
	}
}

func TestGamer_IsTurn(t *testing.T) {
	var c core.GameCore
	c.InitTurnGamerId(0)
	if !c.IsTurn(0) {
		t.Error("((")
	}
}

func TestGamer_Move(t *testing.T) {
	var c core.GameCore
	c.InitField(core.NewStandard8x8Field())
	var gamer [2]Gamer
	gamer[0] = Gamer{0, &c}
	gamer[1] = Gamer{1, &c}
	if gamer[0].Move(core.Coordinate{2, 1}, []core.Coordinate{{3, 0}}) {
		t.Error()
	}
	if gamer[1].Move(core.Coordinate{2, 0}, []core.Coordinate{{3, 1}}) {
		t.Error(c)
	}
	if gamer[1].Move(core.Coordinate{2, 1}, []core.Coordinate{{3, 0}}) {
		t.Error()
	}
	if gamer[1].Move(core.Coordinate{2, 0}, []core.Coordinate{{3, 1}}) {
		t.Error()
	}
	if !gamer[0].Move(
		core.Coordinate{2, 0}, []core.Coordinate{{3, 1}},
	) {
		t.Error()
	}
	if c.Move(core.Coordinate{2, 2}, []core.Coordinate{{3, 3}}, 0) {
		t.Error()
	}
	if c.Move(core.Coordinate{2, 2}, []core.Coordinate{{3, 3}}, 1) {
		t.Error()
	}
	if !c.Move(core.Coordinate{5, 1}, []core.Coordinate{{4, 2}}, 1) {
		t.Error()
	}
}

func TestGamer_GetWinner(t *testing.T) {

}

func TestGamer_InitSave(t *testing.T) {

}
