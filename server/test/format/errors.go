package format

import "fmt"

func ErrorInt(expected, got int) string {
	return fmt.Sprintf("expected %d, but got %d", expected, got)
}

func ErrorString(expected, got string) string {
	return fmt.Sprintf("expected %s, but got %s", expected, got)
}
