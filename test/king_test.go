package test

import (
	"chekers/core"
	"testing"
)

func TestKing_GetOwnerId(t *testing.T) {
	king := core.King{0}

	if king.GetOwnerId() != 0 {
		t.Error()
	}
}

func TestKing_IsMoveOne(t *testing.T) {
	field := createTestField()
	king := core.King{0}
	from := core.Coordinate{3, 3}
	field.Put(from, king)

	test := func(to core.Coordinate, returnValue bool) {
		ok, _ := king.IsMoveOne(&field, from, to)
		if ok != returnValue {
			t.Error()
		}
	}

	test(core.Coordinate{0, 0}, true)
	test(core.Coordinate{2, 2}, true)
	test(core.Coordinate{7, 7}, true)
	test(core.Coordinate{0, 6}, true)
	test(core.Coordinate{7, 7}, true)
	test(core.Coordinate{6, 0}, true)

	test(core.Coordinate{3, 3}, false)
	test(core.Coordinate{-1, -1}, false)
	test(core.Coordinate{8, 8}, false)
	test(core.Coordinate{6, 3}, false)
	test(core.Coordinate{8, 1}, false)
}

func TestKing_Move_WithoutFood(t *testing.T) {
	field := createTestField()
	king := core.King{0}
	from := core.Coordinate{3, 3}
	field.Put(from, king)

	test := func(to core.Coordinate, returnValue bool) {
		isMoved, where := king.Move(&field, from, []core.Coordinate{to})
		if isMoved != returnValue {
			t.Error(to)
		}
		if isMoved == true && where != to {
			t.Error(to)
		}
		if isMoved == false && where != from {
			t.Error(to)
		}
		if isMoved {
			if field.At(to) != king {
				t.Error(to)
			}
			if field.At(from) != nil {
				t.Error(to)
			}
			field.Move(to, from)
		}
		if !isMoved && (field.At(from) != king || field.At(to) != nil) {
			if to != from {
				t.Error(to)
			}
			if field.At(from) != king {
				t.Error(to)
			}
		}
	}

	test(core.Coordinate{0, 0}, true)
	test(core.Coordinate{2, 2}, true)
	test(core.Coordinate{7, 7}, true)
	test(core.Coordinate{0, 6}, true)
	test(core.Coordinate{7, 7}, true)
	test(core.Coordinate{6, 0}, true)

	test(core.Coordinate{3, 3}, false)
	test(core.Coordinate{-1, -1}, false)
	test(core.Coordinate{8, 8}, false)
	test(core.Coordinate{6, 3}, false)
	test(core.Coordinate{8, 1}, false)
}

func testFigure_Move(t *testing.T, figure core.Figure, fininshFigure core.Figure,
	checkField core.Field, field core.Field, from core.Coordinate, to []core.Coordinate,
	empty map[core.Coordinate]bool, _isMoved bool, _finish core.Coordinate,
	messeage string) {

	field.Put(from, figure)
	isMoved, finish := figure.Move(&field, from, to)
	if isMoved != _isMoved {
		t.Error()
	}
	if finish != _finish {
		t.Error(messeage)
	}
	for ptr, _ := range checkField.Figures {
		if field.At(ptr) == fininshFigure {
			if ptr != finish {
				t.Error(messeage)
			}
			continue
		}
		if empty[ptr] && field.At(ptr) != nil {
			t.Error(messeage)
		} else if !empty[ptr] && field.At(ptr) != checkField.At(ptr) {
			t.Error(messeage)
		}
	}
	if len(field.Bin) != len(empty) {
		t.Error(messeage)
	}
}

func TestKing_Move_WithFood(t *testing.T) {
	var grandField core.Field
	grandField.Init()
	grandField.BordersLeft = core.Coordinate{0, 0}
	grandField.BordersRight = core.Coordinate{8, 8}

	for x := 1; x < 8; x += 2 {
		for y := 1; y < 8; y += 2 {
			grandField.Put(core.Coordinate{x, y}, TestFigure{1})
		}
	}

	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{6, 6}},
		map[core.Coordinate]bool{{5, 5}: true},
		true, core.Coordinate{6, 6},
		"1")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{6, 2}},
		map[core.Coordinate]bool{{5, 3}: true},
		true, core.Coordinate{6, 2},
		"2")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{2, 6}},
		map[core.Coordinate]bool{{3, 5}: true},
		true, core.Coordinate{2, 6},
		"3")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{2, 2}},
		map[core.Coordinate]bool{{3, 3}: true},
		true, core.Coordinate{2, 2},
		"4")

	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(),
		core.Coordinate{8, 8},
		[]core.Coordinate{{6, 6}, {8, 4}, {6, 2}, {4, 0},
			{2, 2}, {0, 4}, {2, 6}, {4, 8},
			{6, 6}, {4, 4}},
		map[core.Coordinate]bool{{7, 7}: true, {7, 5}: true,
			{7, 3}: true, {5, 1}: true, {3, 1}: true,
			{1, 3}: true, {1, 5}: true, {3, 7}: true,
			{5, 7}: true, {5, 5}: true},
		true, core.Coordinate{4, 4},
		"5")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(),
		core.Coordinate{8, 8},
		[]core.Coordinate{{6, 6}, {4, 8}, {2, 6}, {0, 4},
			{2, 2}, {4, 0}, {6, 2}, {8, 4},
			{6, 6}, {4, 4}},
		map[core.Coordinate]bool{{7, 7}: true, {7, 5}: true,
			{7, 3}: true, {5, 1}: true, {3, 1}: true,
			{1, 3}: true, {1, 5}: true, {3, 7}: true,
			{5, 7}: true, {5, 5}: true},
		true, core.Coordinate{4, 4},
		"6")

	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{5, 3}, {6, 2}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"7")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{3, 5}, {2, 6}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"8")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{3, 3}, {2, 2}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"9")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{5, 5}, {6, 6}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"10")

	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{7, 1}, {8, 0}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"11")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{1, 7}, {0, 8}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"12")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{7, 7}, {8, 8}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"13")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{1, 1}, {0, 0}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"14")

	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{8, 0}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"15")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{0, 8}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"16")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{8, 8}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"17")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{0, 0}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"18")

	testFigure_Move(t,
		core.King{0}, core.King{0},
		grandField, grandField.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{6, 2}, {2, 6}, {8, 0}, {0, 8}},
		map[core.Coordinate]bool{{3, 5}: true, {5, 3}: true,
			{1, 7}: true, {7, 1}: true},
		true, core.Coordinate{0, 8},
		"19")

	field := createTestField()
	field.Put(core.Coordinate{3, 3}, TestFigure{1})

	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{2, 2}},
		map[core.Coordinate]bool{},
		true, core.Coordinate{2, 2},
		"20")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{3, 3}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{0, 0},
		"21")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{4, 4}},
		map[core.Coordinate]bool{{3, 3}: true},
		true, core.Coordinate{4, 4},
		"22")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{5, 5}},
		map[core.Coordinate]bool{{3, 3}: true},
		true, core.Coordinate{5, 5},
		"23")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{6, 6}},
		map[core.Coordinate]bool{{3, 3}: true},
		true, core.Coordinate{6, 6},
		"24")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{7, 7}},
		map[core.Coordinate]bool{{3, 3}: true},
		true, core.Coordinate{7, 7},
		"25")

	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{6, 6}, {7, 5}},
		map[core.Coordinate]bool{{3, 3}: true},
		true, core.Coordinate{6, 6},
		"phantom move after eat")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{2, 2}, {3, 1}},
		map[core.Coordinate]bool{},
		true, core.Coordinate{2, 2},
		"phantom move after move without food")

	field = createTestField()
	field.Put(core.Coordinate{3, 3}, TestFigure{0})
	field.Put(core.Coordinate{5, 3}, TestFigure{0})

	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{4, 4}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{0, 0},
		"eat friends")
	testFigure_Move(t,
		core.King{0}, core.King{0},
		field, field.GetCopy(), core.Coordinate{0, 0},
		[]core.Coordinate{{4, 4}, {6, 2}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{0, 0},
		"eat friends")
}

func TestKing_GetAvailableMoves(t *testing.T) {
	field := createTestField()
	field.Put(core.Coordinate{3, 3}, core.King{0})
	field.Put(core.Coordinate{6, 6}, TestFigure{1})
	field.Put(core.Coordinate{2, 2}, TestFigure{1})
	field.Put(core.Coordinate{4, 2}, TestFigure{0})

	figure := field.At(core.Coordinate{3, 3})
	test := func(
		get func(field2 *core.Field,
			coordinate core.Coordinate) []core.Coordinate,
		realMoves map[core.Coordinate]bool) {
		moves := get(&field, core.Coordinate{3, 3})

		for _, ptr := range moves {
			if !realMoves[ptr] {
				t.Error("in moves extra move:", ptr)
			}
			delete(realMoves, ptr)
		}
		for key, _ := range realMoves {
			t.Error("in moves expected:", key, ", but didn't state")
		}
	}

	realMoves := map[core.Coordinate]bool{{4, 4}: true, {5, 5}: true,
		{7, 7}: true, {2, 4}: true, {1, 5}: true,
		{0, 6}: true, {1, 1}: true, {0, 0}: true}
	test(figure.GetAvailableMoves, realMoves)
	realMoves = map[core.Coordinate]bool{{7, 7}: true, {1, 1}: true,
		{0, 0}: true}
	test(figure.GetAvailableMovesToEat, realMoves)
}
