package test

import (
	"chekers/core"
	"testing"
)

func getTestCore() core.GameCore {
	var Core core.GameCore
	field := getTestField()
	field.Put(core.Coordinate{4, 4}, core.Checker{1})
	field.Put(core.Coordinate{1, 1}, core.Checker{0})
	Core.InitTurnGamerId(1)
	Core.InitField(field)
	return Core
}

func TestGameCore_GetField(t *testing.T) {
	Core := getTestCore()
	if Core.GetField().Figures == nil {
		t.Error()
	}
}

func TestGameCore_IsTurn(t *testing.T) {
	Core := getTestCore()
	if !Core.IsTurn(1) {
		t.Error()
	}
}

func TestGameCore_Move(t *testing.T) {
	Core := getTestCore()

	if Core.Move(core.Coordinate{1, 1}, []core.Coordinate{{2, 0}}, 1) {
		t.Error()
	}
	if !Core.IsTurn(1) {
		t.Error()
	}

	if Core.Move(core.Coordinate{1, 1}, []core.Coordinate{{2, 0}}, 0) {
		t.Error()
	}
	if !Core.IsTurn(1) {
		t.Error()
	}

	if Core.Move(core.Coordinate{4, 4}, []core.Coordinate{{5, 5}}, 1) {
		t.Error()
	}
	if !Core.Move(core.Coordinate{4, 4}, []core.Coordinate{{3, 3}}, 1) {
		t.Error()
	}
	if Core.IsTurn(1) {
		t.Error()
	}

	if !Core.Move(core.Coordinate{1, 1}, []core.Coordinate{{2, 2}}, 0) {
		t.Error()
	}
}
