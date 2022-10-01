package core

import (
	"testing"
)

func TestKing_GetOwnerId(t *testing.T) {
	king := King{0}

	if king.GetOwnerId() != 0 {
		t.Error()
	}
}

func TestKing_IsMoveOne(t *testing.T) {
	field := NewTestField()
	king := King{0}
	from := Coordinate{3, 3}
	field.Put(from, king)

	test := func(to Coordinate, returnValue bool) {
		ok, _ := king.IsMoveOne(&field, from, to)
		if ok != returnValue {
			t.Error()
		}
	}

	test(Coordinate{0, 0}, true)
	test(Coordinate{2, 2}, true)
	test(Coordinate{7, 7}, true)
	test(Coordinate{0, 6}, true)
	test(Coordinate{7, 7}, true)
	test(Coordinate{6, 0}, true)

	test(Coordinate{3, 3}, false)
	test(Coordinate{-1, -1}, false)
	test(Coordinate{8, 8}, false)
	test(Coordinate{6, 3}, false)
	test(Coordinate{8, 1}, false)
}

func TestKing_Move_WithoutFood(t *testing.T) {
	field := NewTestField()
	king := King{0}
	from := Coordinate{3, 3}
	field.Put(from, king)

	test := func(to Coordinate, returnValue bool) {
		isMoved, where := king.Move(&field, from, []Coordinate{to})
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

	test(Coordinate{0, 0}, true)
	test(Coordinate{2, 2}, true)
	test(Coordinate{7, 7}, true)
	test(Coordinate{0, 6}, true)
	test(Coordinate{7, 7}, true)
	test(Coordinate{6, 0}, true)

	test(Coordinate{3, 3}, false)
	test(Coordinate{-1, -1}, false)
	test(Coordinate{8, 8}, false)
	test(Coordinate{6, 3}, false)
	test(Coordinate{8, 1}, false)
}

func testFigure_Move(
	t *testing.T, figure Figure, fininshFigure Figure,
	checkField Field, field Field, from Coordinate, to []Coordinate,
	empty map[Coordinate]bool, _isMoved bool, _finish Coordinate,
	messeage string,
) {

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
	grandField := NewField()
	grandField.BordersLeft = Coordinate{0, 0}
	grandField.BordersRight = Coordinate{8, 8}

	for x := 1; x < 8; x += 2 {
		for y := 1; y < 8; y += 2 {
			grandField.Put(Coordinate{x, y}, TestFigure{1})
		}
	}

	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{6, 6}},
		map[Coordinate]bool{{5, 5}: true},
		true, Coordinate{6, 6},
		"1",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{6, 2}},
		map[Coordinate]bool{{5, 3}: true},
		true, Coordinate{6, 2},
		"2",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{2, 6}},
		map[Coordinate]bool{{3, 5}: true},
		true, Coordinate{2, 6},
		"3",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{2, 2}},
		map[Coordinate]bool{{3, 3}: true},
		true, Coordinate{2, 2},
		"4",
	)

	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(),
		Coordinate{8, 8},
		[]Coordinate{
			{6, 6}, {8, 4}, {6, 2}, {4, 0},
			{2, 2}, {0, 4}, {2, 6}, {4, 8},
			{6, 6}, {4, 4},
		},
		map[Coordinate]bool{
			{7, 7}: true, {7, 5}: true,
			{7, 3}: true, {5, 1}: true, {3, 1}: true,
			{1, 3}: true, {1, 5}: true, {3, 7}: true,
			{5, 7}: true, {5, 5}: true,
		},
		true, Coordinate{4, 4},
		"5",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(),
		Coordinate{8, 8},
		[]Coordinate{
			{6, 6}, {4, 8}, {2, 6}, {0, 4},
			{2, 2}, {4, 0}, {6, 2}, {8, 4},
			{6, 6}, {4, 4},
		},
		map[Coordinate]bool{
			{7, 7}: true, {7, 5}: true,
			{7, 3}: true, {5, 1}: true, {3, 1}: true,
			{1, 3}: true, {1, 5}: true, {3, 7}: true,
			{5, 7}: true, {5, 5}: true,
		},
		true, Coordinate{4, 4},
		"6",
	)

	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{5, 3}, {6, 2}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"7",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{3, 5}, {2, 6}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"8",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{3, 3}, {2, 2}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"9",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{5, 5}, {6, 6}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"10",
	)

	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{7, 1}, {8, 0}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"11",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{1, 7}, {0, 8}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"12",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{7, 7}, {8, 8}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"13",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{1, 1}, {0, 0}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"14",
	)

	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{8, 0}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"15",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{0, 8}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"16",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{8, 8}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"17",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{0, 0}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"18",
	)

	testFigure_Move(
		t,
		King{0}, King{0},
		grandField, grandField.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{6, 2}, {2, 6}, {8, 0}, {0, 8}},
		map[Coordinate]bool{
			{3, 5}: true, {5, 3}: true,
			{1, 7}: true, {7, 1}: true,
		},
		true, Coordinate{0, 8},
		"19",
	)

	field := NewTestField()
	field.Put(Coordinate{3, 3}, TestFigure{1})

	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{2, 2}},
		map[Coordinate]bool{},
		true, Coordinate{2, 2},
		"20",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{3, 3}},
		map[Coordinate]bool{},
		false, Coordinate{0, 0},
		"21",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{4, 4}},
		map[Coordinate]bool{{3, 3}: true},
		true, Coordinate{4, 4},
		"22",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{5, 5}},
		map[Coordinate]bool{{3, 3}: true},
		true, Coordinate{5, 5},
		"23",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{6, 6}},
		map[Coordinate]bool{{3, 3}: true},
		true, Coordinate{6, 6},
		"24",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{7, 7}},
		map[Coordinate]bool{{3, 3}: true},
		true, Coordinate{7, 7},
		"25",
	)

	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{6, 6}, {7, 5}},
		map[Coordinate]bool{{3, 3}: true},
		true, Coordinate{6, 6},
		"phantom move after eat",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{2, 2}, {3, 1}},
		map[Coordinate]bool{},
		true, Coordinate{2, 2},
		"phantom move after move without food",
	)

	field = NewTestField()
	field.Put(Coordinate{3, 3}, TestFigure{0})
	field.Put(Coordinate{5, 3}, TestFigure{0})

	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{4, 4}},
		map[Coordinate]bool{},
		false, Coordinate{0, 0},
		"eat friends",
	)
	testFigure_Move(
		t,
		King{0}, King{0},
		field, field.GetCopy(), Coordinate{0, 0},
		[]Coordinate{{4, 4}, {6, 2}},
		map[Coordinate]bool{},
		false, Coordinate{0, 0},
		"eat friends",
	)
}

func TestKing_GetAvailableMoves(t *testing.T) {
	field := NewTestField()
	field.Put(Coordinate{3, 3}, King{0})
	field.Put(Coordinate{6, 6}, TestFigure{1})
	field.Put(Coordinate{2, 2}, TestFigure{1})
	field.Put(Coordinate{4, 2}, TestFigure{0})

	figure := field.At(Coordinate{3, 3})
	test := func(
		get func(
			field2 *Field,
			coordinate Coordinate,
		) []Coordinate,
		realMoves map[Coordinate]bool,
	) {
		moves := get(&field, Coordinate{3, 3})

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

	realMoves := map[Coordinate]bool{
		{4, 4}: true, {5, 5}: true,
		{7, 7}: true, {2, 4}: true, {1, 5}: true,
		{0, 6}: true, {1, 1}: true, {0, 0}: true,
	}
	test(figure.GetAvailableMoves, realMoves)
	realMoves = map[Coordinate]bool{
		{7, 7}: true, {1, 1}: true,
		{0, 0}: true,
	}
	test(figure.GetAvailableMovesToEat, realMoves)
}
