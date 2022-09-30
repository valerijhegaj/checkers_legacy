package core

import (
	"reflect"
	"testing"
)

func TestField_InBorders(t *testing.T) {
	var field Field
	field.BordersRight.X = 7
	field.BordersRight.Y = 7
	field.BordersLeft.X = 1
	field.BordersLeft.Y = 1

	if !field.InBorders(Coordinate{1, 1}) {
		t.Error("got false, expected true")
	}
	if !field.InBorders(Coordinate{7, 7}) {
		t.Error("got false, expected true")
	}
	if !field.InBorders(Coordinate{7, 1}) {
		t.Error("got false, expected true")
	}
	if !field.InBorders(Coordinate{1, 7}) {
		t.Error("got false, expected true")
	}

	if field.InBorders(Coordinate{-1, 0}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(Coordinate{-1, -1}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(Coordinate{0, -1}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(Coordinate{8, 7}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(Coordinate{8, 8}) {
		t.Error("got true, expected false")
	}
	if field.InBorders(Coordinate{7, 8}) {
		t.Error("got true, expected false")
	}
}

func NewTestField() Field {
	field := NewField()
	field.BordersRight = Coordinate{7, 7}
	field.BordersLeft = Coordinate{0, 0}

	return field
}

func TestField_Put(t *testing.T) {
	field := NewTestField()
	for i := 0; i < 3; i++ {
		field.Put(Coordinate{i, i}, TestFigure{1})
		if len(field.Figures) != i+1 {
			t.Error("can't put")
		}
	}
}

func TestField_At(t *testing.T) {
	field := NewTestField()

	field.Put(Coordinate{1, 1}, TestFigure{1})
	field.Put(Coordinate{2, 2}, TestFigure{0})

	figure := field.At(Coordinate{1, 1})
	if figure == nil {
		t.Error("don't extract correctly")
	} else if figure.GetOwnerId() != 1 {
		t.Error("incorrect ownerId of extracted")
	}
	figure = field.At(Coordinate{2, 2})
	if figure == nil {
		t.Error("don't extract correctly")
	} else if figure.GetOwnerId() != 0 {
		t.Error("incorrect ownerId of extracted")
	}
	figure = field.At(Coordinate{1, 2})
	if figure != nil {
		t.Error("expected nil")
	}
	figure = field.At(Coordinate{0, 0})
	if figure != nil {
		t.Error("expected nil")
	}
}

func TestField_IsAvailable(t *testing.T) {
	field := NewTestField()
	if !field.IsAvailable(Coordinate{0, 0}) {
		t.Error()
	}
	if field.IsAvailable(Coordinate{-1, 0}) {
		t.Error()
	}
	field.Put(Coordinate{1, 1}, TestFigure{1})
	field.Put(Coordinate{2, 2}, TestFigure{1})

	if field.IsAvailable(Coordinate{1, 1}) {
		t.Error()
	}
	if field.IsAvailable(Coordinate{2, 2}) {
		t.Error()
	}
	if !field.IsAvailable(Coordinate{2, 1}) {
		t.Error()
	}
	if field.IsAvailable(Coordinate{8, 7}) {
		t.Error()
	}
	if !field.IsAvailable(Coordinate{7, 7}) {
		t.Error()
	}
}

func TestField_Move(t *testing.T) {
	field := NewTestField()
	field.Put(Coordinate{0, 0}, TestFigure{0})
	field.Move(Coordinate{0, 0}, Coordinate{1, 1})

	if field.At(Coordinate{0, 0}) != nil {
		t.Error()
	}
	if field.At(Coordinate{1, 1}) == nil {
		t.Error()
	} else if field.At(Coordinate{1, 1}).GetOwnerId() != 0 {
		t.Error()
	} else if reflect.TypeOf(field.At(Coordinate{1, 1})) != reflect.TypeOf(TestFigure{}) {
		t.Error()
	}

	field.Move(Coordinate{3, 3}, Coordinate{2, 2})
	if field.At(Coordinate{3, 3}) != nil {
		t.Error()
	}
	if field.At(Coordinate{2, 2}) != nil {
		t.Error()
	}
}

func TestField_Remove(t *testing.T) {
	field := NewTestField()

	field.Put(Coordinate{0, 0}, TestFigure{0})
	field.Remove(Coordinate{0, 0})

	if len(field.Bin) != 1 {
		t.Error()
	}
	if field.Bin[0].GetOwnerId() != 0 {
		t.Error()
	}
	if reflect.TypeOf(field.Bin[0]) != reflect.TypeOf(TestFigure{0}) {
		t.Error()
	}
	if field.At(Coordinate{0, 0}) != nil {
		t.Error()
	}
}

func TestField_RemoveWithOutBin(t *testing.T) {
	field := NewTestField()

	field.Put(Coordinate{0, 0}, TestFigure{0})
	field.RemoveWithOutBin(Coordinate{0, 0})

	if len(field.Bin) != 0 {
		t.Error()
	}
	if field.At(Coordinate{0, 0}) != nil {
		t.Error()
	}
}

func TestField_GetCopy(t *testing.T) {
	field := NewTestField()
	field.Put(Coordinate{0, 0}, TestFigure{0})
	copy := field.GetCopy()
	field.Remove(Coordinate{0, 0})
	field.Put(Coordinate{0, 1}, TestFigure{1})

	if len(copy.Bin) == 1 {
		t.Error()
	}
	if copy.At(Coordinate{0, 0}) == nil {
		t.Error()
	}
	if copy.At(Coordinate{0, 1}) != nil {
		t.Error()
	}
}
