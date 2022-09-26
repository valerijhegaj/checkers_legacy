package test

import (
	"chekers/core"
	"testing"
)

func TestChecker_GetOwnerId(t *testing.T) {
	checker := core.Checker{0}

	if checker.GetOwnerId() != 0 {
		t.Error()
	}
}

func TestChecker_IsMoveOne0(t *testing.T) {
	field := getTestField()
	checker := core.Checker{0}
	from := core.Coordinate{3, 3}
	field.Put(from, checker)

	test := func(to core.Coordinate, returnValue bool) {
		ok, _ := checker.IsMoveOne(&field, from, to)
		if ok != returnValue {
			t.Error()
		}
	}

	true_1 := core.Coordinate{4, 2}
	true_2 := core.Coordinate{4, 4}

	for x := 1; x < 6; x++ {
		for y := 1; y < 6; y++ {
			if (x == true_1.X && y == true_1.Y) || (x == true_2.X && y == true_2.Y) {
				test(core.Coordinate{x, y}, true)
				continue
			}
			test(core.Coordinate{x, y}, false)
		}
	}

	field.Put(true_1, core.Checker{0})
	field.Put(true_2, core.Checker{1})

	true_1 = core.Coordinate{5, 5}

	for x := 1; x < 6; x++ {
		for y := 1; y < 6; y++ {
			if x == true_1.X && y == true_1.Y {
				test(core.Coordinate{x, y}, true)
				continue
			}
			test(core.Coordinate{x, y}, false)
		}
	}
}

func TestChecker_IsMoveOne1(t *testing.T) {
	field := getTestField()
	checker := core.Checker{1}
	from := core.Coordinate{3, 3}
	field.Put(from, checker)

	test := func(to core.Coordinate, returnValue bool) {
		ok, _ := checker.IsMoveOne(&field, from, to)
		if ok != returnValue {
			t.Error()
		}
	}

	true_1 := core.Coordinate{2, 2}
	true_2 := core.Coordinate{2, 4}

	for x := 1; x < 6; x++ {
		for y := 1; y < 6; y++ {
			if (x == true_1.X && y == true_1.Y) || (x == true_2.X && y == true_2.Y) {
				test(core.Coordinate{x, y}, true)
				continue
			}
			test(core.Coordinate{x, y}, false)
		}
	}

	field.Put(true_1, core.Checker{0})
	field.Put(true_2, core.Checker{1})

	true_1 = core.Coordinate{1, 1}

	for x := 1; x < 6; x++ {
		for y := 1; y < 6; y++ {
			if x == true_1.X && y == true_1.Y {
				test(core.Coordinate{x, y}, true)
				continue
			}
			test(core.Coordinate{x, y}, false)
		}
	}
}

func TestChecker_Move(t *testing.T) {
	var grandField0, grandField1 core.Field

	grandField0.Init()
	grandField0.BordersLeft = core.Coordinate{-1, -1}
	grandField0.BordersRight = core.Coordinate{9, 9}

	grandField1.Init()
	grandField1.BordersLeft = core.Coordinate{-1, -1}
	grandField1.BordersRight = core.Coordinate{9, 9}

	for x := 1; x < 8; x += 2 {
		for y := 1; y < 8; y += 2 {
			grandField0.Put(core.Coordinate{x, y}, Test_figure{0})
		}
	}

	for x := 1; x < 8; x += 2 {
		for y := 1; y < 8; y += 2 {
			grandField1.Put(core.Coordinate{x, y}, Test_figure{1})
		}
	}

	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{6, 6}},
		map[core.Coordinate]bool{{5, 5}: true},
		true, core.Coordinate{6, 6},
		"1")
	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{6, 2}},
		map[core.Coordinate]bool{{5, 3}: true},
		true, core.Coordinate{6, 2},
		"2")
	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{2, 6}},
		map[core.Coordinate]bool{{3, 5}: true},
		true, core.Coordinate{2, 6},
		"3")
	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{2, 2}},
		map[core.Coordinate]bool{{3, 3}: true},
		true, core.Coordinate{2, 2},
		"4")

	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(),
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
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(),
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
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{5, 3}, {6, 2}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"7")
	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{3, 5}, {2, 6}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"8")
	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{3, 3}, {2, 2}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"9")
	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{4, 4},
		[]core.Coordinate{{5, 5}, {6, 6}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{4, 4},
		"10")

	testFigure_Move(t,
		core.Checker{0}, core.King{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{8, 4},
		[]core.Coordinate{{9, 5}},
		map[core.Coordinate]bool{},
		true, core.Coordinate{9, 5},
		"turninig to king")
	testFigure_Move(t,
		core.Checker{1}, core.Checker{1},
		grandField1, grandField1.GetCopy(), core.Coordinate{8, 4},
		[]core.Coordinate{{9, 5}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{8, 4},
		"turninig to king")
	testFigure_Move(t,
		core.Checker{1}, core.King{1},
		grandField1, grandField1.GetCopy(), core.Coordinate{0, 4},
		[]core.Coordinate{{-1, 3}},
		map[core.Coordinate]bool{},
		true, core.Coordinate{-1, 3},
		"turninig to king")
	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{0, 4},
		[]core.Coordinate{{-1, 3}},
		map[core.Coordinate]bool{},
		false, core.Coordinate{0, 4},
		"turninig to king")

	testFigure_Move(t,
		core.Checker{0}, core.King{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{8, 4},
		[]core.Coordinate{{9, 5}, {8, 4}},
		map[core.Coordinate]bool{},
		true, core.Coordinate{9, 5},
		"fantom move after turninig to king")
	testFigure_Move(t,
		core.Checker{1}, core.King{1},
		grandField1, grandField1.GetCopy(), core.Coordinate{0, 4},
		[]core.Coordinate{{-1, 3}, {8, 4}},
		map[core.Coordinate]bool{},
		true, core.Coordinate{-1, 3},
		"fantom move after turninig to king")

	grandField1.BordersLeft = core.Coordinate{0, 0}
	grandField1.BordersRight = core.Coordinate{8, 8}

	testFigure_Move(t,
		core.Checker{0}, core.King{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{6, 6},
		[]core.Coordinate{{8, 4}},
		map[core.Coordinate]bool{{7, 5}: true},
		true, core.Coordinate{8, 4},
		"turninig to king after eat")
	testFigure_Move(t,
		core.Checker{1}, core.King{1},
		grandField0, grandField0.GetCopy(), core.Coordinate{2, 6},
		[]core.Coordinate{{0, 4}},
		map[core.Coordinate]bool{{1, 5}: true},
		true, core.Coordinate{0, 4},
		"turninig to king after eat")

	testFigure_Move(t,
		core.Checker{0}, core.King{0},
		grandField1, grandField1.GetCopy(), core.Coordinate{6, 6},
		[]core.Coordinate{{8, 4}, {6, 2}},
		map[core.Coordinate]bool{{7, 5}: true, {7, 3}: true},
		true, core.Coordinate{6, 2},
		"eat after turninig to king")
	testFigure_Move(t,
		core.Checker{1}, core.King{1},
		grandField0, grandField0.GetCopy(), core.Coordinate{2, 6},
		[]core.Coordinate{{0, 4}, {2, 2}},
		map[core.Coordinate]bool{{1, 5}: true, {1, 3}: true},
		true, core.Coordinate{2, 2},
		"eat after turninig to king")

	testFigure_Move(t,
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(),
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
		core.Checker{0}, core.Checker{0},
		grandField1, grandField1.GetCopy(),
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
}
