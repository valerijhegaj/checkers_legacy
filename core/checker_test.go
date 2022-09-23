package core

import "testing"

func TestChecker_GetOwnerId(t *testing.T) {
	checker := Checker{0}

	if checker.GetOwnerId() != 0 {
		t.Error()
	}
}

func TestChecker_IsMoveOne0(t *testing.T) {
	field := getTestField()
	checker := Checker{0}
	from := Coordinate{3, 3}
	field.Put(from, checker)

	test := func(to Coordinate, returnValue bool) {
		if checker.IsMoveOne(&field, from, to) != returnValue {
			t.Error()
		}
	}

	true_1 := Coordinate{4, 2}
	true_2 := Coordinate{4, 4}

	for x := 1; x < 6; x++ {
		for y := 1; y < 6; y++ {
			if (x == true_1.X && y == true_1.Y) || (x == true_2.X && y == true_2.Y) {
				test(Coordinate{x, y}, true)
				continue
			}
			test(Coordinate{x, y}, false)
		}
	}

	field.Put(true_1, Checker{0})
	field.Put(true_2, Checker{1})

	true_1 = Coordinate{5, 5}

	for x := 1; x < 6; x++ {
		for y := 1; y < 6; y++ {
			if x == true_1.X && y == true_1.Y {
				test(Coordinate{x, y}, true)
				continue
			}
			test(Coordinate{x, y}, false)
		}
	}
}

func TestChecker_IsMoveOne1(t *testing.T) {
	field := getTestField()
	checker := Checker{1}
	from := Coordinate{3, 3}
	field.Put(from, checker)

	test := func(to Coordinate, returnValue bool) {
		if checker.IsMoveOne(&field, from, to) != returnValue {
			t.Error()
		}
	}

	true_1 := Coordinate{2, 2}
	true_2 := Coordinate{2, 4}

	for x := 1; x < 6; x++ {
		for y := 1; y < 6; y++ {
			if (x == true_1.X && y == true_1.Y) || (x == true_2.X && y == true_2.Y) {
				test(Coordinate{x, y}, true)
				continue
			}
			test(Coordinate{x, y}, false)
		}
	}

	field.Put(true_1, Checker{0})
	field.Put(true_2, Checker{1})

	true_1 = Coordinate{1, 1}

	for x := 1; x < 6; x++ {
		for y := 1; y < 6; y++ {
			if x == true_1.X && y == true_1.Y {
				test(Coordinate{x, y}, true)
				continue
			}
			test(Coordinate{x, y}, false)
		}
	}
}

func TestChecker_Move(t *testing.T) {
	var grandField0, grandField1 Field

	grandField0.Init()
	grandField0.BordersLeft = Coordinate{-1, -1}
	grandField0.BordersRight = Coordinate{9, 9}

	grandField1.Init()
	grandField1.BordersLeft = Coordinate{-1, -1}
	grandField1.BordersRight = Coordinate{9, 9}

	for x := 1; x < 8; x += 2 {
		for y := 1; y < 8; y += 2 {
			grandField0.Put(Coordinate{x, y}, test_figure{0})
		}
	}

	for x := 1; x < 8; x += 2 {
		for y := 1; y < 8; y += 2 {
			grandField1.Put(Coordinate{x, y}, test_figure{1})
		}
	}

	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{6, 6}},
		map[Coordinate]bool{{5, 5}: true},
		true, Coordinate{6, 6},
		"1")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{6, 2}},
		map[Coordinate]bool{{5, 3}: true},
		true, Coordinate{6, 2},
		"2")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{2, 6}},
		map[Coordinate]bool{{3, 5}: true},
		true, Coordinate{2, 6},
		"3")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{2, 2}},
		map[Coordinate]bool{{3, 3}: true},
		true, Coordinate{2, 2},
		"4")

	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(),
		Coordinate{8, 8},
		[]Coordinate{{6, 6}, {8, 4}, {6, 2}, {4, 0},
			{2, 2}, {0, 4}, {2, 6}, {4, 8},
			{6, 6}, {4, 4}},
		map[Coordinate]bool{{7, 7}: true, {7, 5}: true,
			{7, 3}: true, {5, 1}: true, {3, 1}: true,
			{1, 3}: true, {1, 5}: true, {3, 7}: true,
			{5, 7}: true, {5, 5}: true},
		true, Coordinate{4, 4},
		"5")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(),
		Coordinate{8, 8},
		[]Coordinate{{6, 6}, {4, 8}, {2, 6}, {0, 4},
			{2, 2}, {4, 0}, {6, 2}, {8, 4},
			{6, 6}, {4, 4}},
		map[Coordinate]bool{{7, 7}: true, {7, 5}: true,
			{7, 3}: true, {5, 1}: true, {3, 1}: true,
			{1, 3}: true, {1, 5}: true, {3, 7}: true,
			{5, 7}: true, {5, 5}: true},
		true, Coordinate{4, 4},
		"6")

	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{5, 3}, {6, 2}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"7")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{3, 5}, {2, 6}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"8")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{3, 3}, {2, 2}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"9")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{4, 4},
		[]Coordinate{{5, 5}, {6, 6}},
		map[Coordinate]bool{},
		false, Coordinate{4, 4},
		"10")

	testFigure_Move(t,
		Checker{0}, King{0},
		grandField1, grandField1.GetCopy(), Coordinate{8, 4},
		[]Coordinate{{9, 5}},
		map[Coordinate]bool{},
		true, Coordinate{9, 5},
		"turninig to king")
	testFigure_Move(t,
		Checker{1}, Checker{1},
		grandField1, grandField1.GetCopy(), Coordinate{8, 4},
		[]Coordinate{{9, 5}},
		map[Coordinate]bool{},
		false, Coordinate{8, 4},
		"turninig to king")
	testFigure_Move(t,
		Checker{1}, King{1},
		grandField1, grandField1.GetCopy(), Coordinate{0, 4},
		[]Coordinate{{-1, 3}},
		map[Coordinate]bool{},
		true, Coordinate{-1, 3},
		"turninig to king")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(), Coordinate{0, 4},
		[]Coordinate{{-1, 3}},
		map[Coordinate]bool{},
		false, Coordinate{0, 4},
		"turninig to king")

	testFigure_Move(t,
		Checker{0}, King{0},
		grandField1, grandField1.GetCopy(), Coordinate{8, 4},
		[]Coordinate{{9, 5}, {8, 4}},
		map[Coordinate]bool{},
		true, Coordinate{9, 5},
		"fantom move after turninig to king")
	testFigure_Move(t,
		Checker{1}, King{1},
		grandField1, grandField1.GetCopy(), Coordinate{0, 4},
		[]Coordinate{{-1, 3}, {8, 4}},
		map[Coordinate]bool{},
		true, Coordinate{-1, 3},
		"fantom move after turninig to king")

	grandField1.BordersLeft = Coordinate{0, 0}
	grandField1.BordersRight = Coordinate{8, 8}

	testFigure_Move(t,
		Checker{0}, King{0},
		grandField1, grandField1.GetCopy(), Coordinate{6, 6},
		[]Coordinate{{8, 4}},
		map[Coordinate]bool{{7, 5}: true},
		true, Coordinate{8, 4},
		"turninig to king after eat")
	testFigure_Move(t,
		Checker{1}, King{1},
		grandField0, grandField0.GetCopy(), Coordinate{2, 6},
		[]Coordinate{{0, 4}},
		map[Coordinate]bool{{1, 5}: true},
		true, Coordinate{0, 4},
		"turninig to king after eat")

	testFigure_Move(t,
		Checker{0}, King{0},
		grandField1, grandField1.GetCopy(), Coordinate{6, 6},
		[]Coordinate{{8, 4}, {6, 2}},
		map[Coordinate]bool{{7, 5}: true, {7, 3}: true},
		true, Coordinate{6, 2},
		"eat after turninig to king")
	testFigure_Move(t,
		Checker{1}, King{1},
		grandField0, grandField0.GetCopy(), Coordinate{2, 6},
		[]Coordinate{{0, 4}, {2, 2}},
		map[Coordinate]bool{{1, 5}: true, {1, 3}: true},
		true, Coordinate{2, 2},
		"eat after turninig to king")

	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(),
		Coordinate{8, 8},
		[]Coordinate{{6, 6}, {8, 4}, {6, 2}, {4, 0},
			{2, 2}, {0, 4}, {2, 6}, {4, 8},
			{6, 6}, {4, 4}},
		map[Coordinate]bool{{7, 7}: true, {7, 5}: true,
			{7, 3}: true, {5, 1}: true, {3, 1}: true,
			{1, 3}: true, {1, 5}: true, {3, 7}: true,
			{5, 7}: true, {5, 5}: true},
		true, Coordinate{4, 4},
		"5")
	testFigure_Move(t,
		Checker{0}, Checker{0},
		grandField1, grandField1.GetCopy(),
		Coordinate{8, 8},
		[]Coordinate{{6, 6}, {4, 8}, {2, 6}, {0, 4},
			{2, 2}, {4, 0}, {6, 2}, {8, 4},
			{6, 6}, {4, 4}},
		map[Coordinate]bool{{7, 7}: true, {7, 5}: true,
			{7, 3}: true, {5, 1}: true, {3, 1}: true,
			{1, 3}: true, {1, 5}: true, {3, 7}: true,
			{5, 7}: true, {5, 5}: true},
		true, Coordinate{4, 4},
		"6")
}
