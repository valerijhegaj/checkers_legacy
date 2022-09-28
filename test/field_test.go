package test

import (
	"chekers/core"
	"reflect"
	"testing"
)

func TestField_InBorders(t *testing.T) {
	var field core.Field
	field.BordersRight.X = 7
	field.BordersRight.Y = 7
	field.BordersLeft.X = 1
	field.BordersLeft.Y = 1

	if !field.InBorders(core.Coordinate{1, 1}) {
		t.Error("got false, expected true")
	}
	if !field.InBorders(core.Coordinate{7, 7}) {
		t.Error("got false, expected true")
	}
	if !field.InBorders(core.Coordinate{7, 1}) {
		t.Error("got false, expected true")
	}
	if !field.InBorders(core.Coordinate{1, 7}) {
		t.Error("got false, expected true")
	}

	if field.InBorders(core.Coordinate{-1, 0}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(core.Coordinate{-1, -1}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(core.Coordinate{0, -1}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(core.Coordinate{8, 7}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(core.Coordinate{8, 8}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(core.Coordinate{7, 8}) {
		t.Error("got true, expected false")
	}
}

func createTestField() core.Field {
	var field core.Field
	field.Init()
	field.BordersRight = core.Coordinate{7, 7}
	field.BordersLeft = core.Coordinate{0, 0}

	return field
}

func TestField_Put(t *testing.T) {
	field := createTestField()
	for i := 0; i < 3; i++ {
		field.Put(core.Coordinate{i, i}, TestFigure{1})
		if len(field.Figures) != i+1 {
			t.Error("can't put")
		}
	}
}

func TestField_At(t *testing.T) {
	field := createTestField()

	field.Put(core.Coordinate{1, 1}, TestFigure{1})
	field.Put(core.Coordinate{2, 2}, TestFigure{0})

	figure := field.At(core.Coordinate{1, 1})
	if figure == nil {
		t.Error("don't extract correctly")
	} else if figure.GetOwnerId() != 1 {
		t.Error("incorrect ownerId of extracted")
	}
	figure = field.At(core.Coordinate{2, 2})
	if figure == nil {
		t.Error("don't extract correctly")
	} else if figure.GetOwnerId() != 0 {
		t.Error("incorrect ownerId of extracted")
	}
	figure = field.At(core.Coordinate{1, 2})
	if figure != nil {
		t.Error("expected nil")
	}
	figure = field.At(core.Coordinate{0, 0})
	if figure != nil {
		t.Error("expected nil")
	}
}

func TestField_IsAvailable(t *testing.T) {
	field := createTestField()
	if !field.IsAvailable(core.Coordinate{0, 0}) {
		t.Error()
	}
	if field.IsAvailable(core.Coordinate{-1, 0}) {
		t.Error()
	}
	field.Put(core.Coordinate{1, 1}, TestFigure{1})
	field.Put(core.Coordinate{2, 2}, TestFigure{1})

	if field.IsAvailable(core.Coordinate{1, 1}) {
		t.Error()
	}
	if field.IsAvailable(core.Coordinate{2, 2}) {
		t.Error()
	}
	if !field.IsAvailable(core.Coordinate{2, 1}) {
		t.Error()
	}
	if field.IsAvailable(core.Coordinate{8, 7}) {
		t.Error()
	}
	if !field.IsAvailable(core.Coordinate{7, 7}) {
		t.Error()
	}
}

func TestField_Move(t *testing.T) {
	field := createTestField()
	field.Put(core.Coordinate{0, 0}, TestFigure{0})
	field.Move(core.Coordinate{0, 0}, core.Coordinate{1, 1})

	if field.At(core.Coordinate{0, 0}) != nil {
		t.Error()
	}
	if field.At(core.Coordinate{1, 1}) == nil {
		t.Error()
	} else if field.At(core.Coordinate{1, 1}).GetOwnerId() != 0 {
		t.Error()
	} else if reflect.TypeOf(field.At(core.Coordinate{1, 1})) != reflect.TypeOf(TestFigure{}) {
		t.Error()
	}

	field.Move(core.Coordinate{3, 3}, core.Coordinate{2, 2})
	if field.At(core.Coordinate{3, 3}) != nil {
		t.Error()
	}
	if field.At(core.Coordinate{2, 2}) != nil {
		t.Error()
	}
}

func TestField_Remove(t *testing.T) {
	field := createTestField()

	field.Put(core.Coordinate{0, 0}, TestFigure{0})
	field.Remove(core.Coordinate{0, 0})

	if len(field.Bin) != 1 {
		t.Error()
	}
	if field.Bin[0].GetOwnerId() != 0 {
		t.Error()
	}
	if reflect.TypeOf(field.Bin[0]) != reflect.TypeOf(TestFigure{0}) {
		t.Error()
	}
	if field.At(core.Coordinate{0, 0}) != nil {
		t.Error()
	}
}

func TestField_RemoveWithOutBin(t *testing.T) {
	field := createTestField()

	field.Put(core.Coordinate{0, 0}, TestFigure{0})
	field.RemoveWithOutBin(core.Coordinate{0, 0})

	if len(field.Bin) != 0 {
		t.Error()
	}
	if field.At(core.Coordinate{0, 0}) != nil {
		t.Error()
	}
}

func TestField_GetCopy(t *testing.T) {
	field := createTestField()
	field.Put(core.Coordinate{0, 0}, TestFigure{0})
	copy := field.GetCopy()
	field.Remove(core.Coordinate{0, 0})
	field.Put(core.Coordinate{0, 1}, TestFigure{1})

	if len(copy.Bin) == 1 {
		t.Error()
	}
	if copy.At(core.Coordinate{0, 0}) == nil {
		t.Error()
	}
	if copy.At(core.Coordinate{0, 1}) != nil {
		t.Error()
	}
}
