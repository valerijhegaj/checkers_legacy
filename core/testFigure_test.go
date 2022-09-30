package core

import (
	"testing"
)

func TestTest_figure_GetOwnerId(t *testing.T) {
	test := func(ownerId int) {
		figure := TestFigure{ownerId}

		got := figure.GetOwnerId()
		if got != ownerId {
			t.Errorf("got %d, wanted %d", got, ownerId)
		}
	}
	test(0)
	test(1)
}
