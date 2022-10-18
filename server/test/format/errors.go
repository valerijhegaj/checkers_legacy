package format

import (
	"fmt"
	"reflect"

	"checkers/core"
)

func ErrorInt(expected, got int) string {
	return fmt.Sprintf("expected %d, but got %d", expected, got)
}

func ErrorString(expected, got string) string {
	return fmt.Sprintf("expected %s, but got %s", expected, got)
}

func ErrorField(expected, got *core.Field) string {
	var ans string
	ans = "\nexpected:\n"
	ans += Field(expected)
	ans += "got:\n"
	ans += Field(got)
	return ans
}

func Field(field *core.Field) string {
	var ans string

	printFigureCorrect := func(x, y int) {
		figure := field.At(core.Coordinate{x, y})
		if figure == nil {
			ans += "_ "
		} else if reflect.TypeOf(figure) == reflect.TypeOf(core.Checker{}) {
			if figure.GetOwnerId() == 1 {
				ans += "\u001B[31m"
			}
			ans += "O \u001B[0m"
		} else {
			if figure.GetOwnerId() == 1 {
				ans += "\u001B[31m"
			}
			ans += "K \u001B[0m"
		}
	}

	for x := field.BordersRight.X; x >= field.BordersLeft.X; x-- {
		ans += fmt.Sprintf("%d ", x+1)
		for y := field.BordersLeft.Y; y <= field.BordersRight.Y; y++ {
			printFigureCorrect(x, y)
		}
		ans += "\n"
	}

	ans += "  "
	for y := field.BordersLeft.Y; y <= field.BordersRight.Y; y++ {
		ans += string(rune('a'+y)) + " "
	}
	ans += "\n"
	return ans
}
