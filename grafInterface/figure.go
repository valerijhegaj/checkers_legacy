//go:generate fyne bundle -o bundle.go ../images

package grafInterface

import (
	"reflect"

	"chekers/core"
	"fyne.io/fyne/v2"
)

func GetResource(figure core.Figure) fyne.Resource {
	if figure.GetOwnerId() == 0 {
		if reflect.TypeOf(figure) == reflect.TypeOf(core.Checker{}) {
			return resourceCheckerGamer0Png
		} else {
			return resourceKingGamer0Png
		}
	} else {
		if reflect.TypeOf(figure) == reflect.TypeOf(core.Checker{}) {
			return resourceCheckerGamer1Png
		} else {
			return resourceKingGamer1Png
		}
	}
}
