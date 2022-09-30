package main

import (
	"chekers/core"
	"chekers/grafInterface"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Checkers")
	var c core.GameCore
	w.Resize(fyne.NewSize(500, 500))
	var interactor grafInterface.Interface
	interactor.Init(&a, &w, &c)
	interactor.Begin(&interactor.MainMenu)
	w.ShowAndRun()
}
