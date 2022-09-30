package test

import (
	"chekers/core"
	"chekers/saveLoad"
	"testing"
)

func getTestCore() core.GameCore {
	var Core core.GameCore
	field := createTestField()
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

func TestChecker_MoveWithFeature1(t *testing.T) {
	var Core core.GameCore
	var save saveLoad.Save
	err := save.Read("../startFields/start_field.json")
	if err != nil {
		t.Error(err.Error())
	}

	Core.InitField(save.Field)
	gamerId := 0
	Core.InitTurnGamerId(gamerId)

	testMove := func(move string, returnValue bool) {
		ok := Core.Move(core.Coordinate{int(move[1] - '1'), int(move[0] - 'a')},
			[]core.Coordinate{{int(move[3] - '1'), int(move[2] - 'a')}}, gamerId)
		if ok {
			gamerId ^= 1
		}
		if ok != returnValue {
			t.Error(move)
		}
	}

	testMove("a3b4", true)
	testMove("b6a5", true)
	testMove("b4c5", true)
	testMove("a5b4", false)
	testMove("d6e5", false)
	testMove("f6e5", false)
	testMove("f6g5", false)
	testMove("h6g5", false)
	testMove("d6b4", true)
	testMove("b2a3", true)
	testMove("c7b6", true)
	testMove("e3d4", false)
	testMove("a3c4", false)
	testMove("a3c5", true)
	testMove("a5b4", false)
	testMove("b6d4", true)
	testMove("e3c5", true)
	testMove("a5b4", true)
	testMove("c5b6", false)
	testMove("c3a5", true)
	testMove("a7b6", true)
	testMove("c5a7", true)
	testMove("b8c7", true)
	testMove("a7b8", true)
	testMove("e7d6", true)
	testMove("d2c3", true)
	testMove("d6e5", true)
	ok := Core.Move(core.Coordinate{7, 1},
		[]core.Coordinate{{5, 3}, {3, 5}}, gamerId)
	if ok {
		gamerId ^= 1
	}
	if ok != true {
		t.Error()
	}
	testMove("d8c7", true)
	testMove("f4d6", false)
	testMove("f4c7", false)
	testMove("f4b6", false)
	testMove("f5b8", false)
	testMove("f4b8", true)
}

func TestChecker_MoveWithFeature2(t *testing.T) {
	var Core core.GameCore
	field := createTestField()
	field.Put(core.Coordinate{4, 1}, core.Checker{1})
	field.Put(core.Coordinate{3, 2}, core.Checker{1})
	field.Put(core.Coordinate{7, 3}, core.Checker{1})
	field.Put(core.Coordinate{1, 4}, core.King{0})
	field.Put(core.Coordinate{0, 1}, core.King{0})
	Core.InitField(field)
	gamerId := 1
	Core.InitTurnGamerId(gamerId)

	testMove := func(move string, returnValue bool) {
		ok := Core.Move(core.Coordinate{int(move[1] - '1'), int(move[0] - 'a')},
			[]core.Coordinate{{int(move[3] - '1'), int(move[2] - 'a')}}, gamerId)
		if ok {
			gamerId ^= 1
		}
		if ok != returnValue {
			t.Error(Core, move)
		}
	}

	testMove("c4d3", true)
	testMove("e2f3", false)
	testMove("b1e4", true)
	testMove("d8c7", true)
	testMove("e4b1", false)
	testMove("e2a6", true)
}
