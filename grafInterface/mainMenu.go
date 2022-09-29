package grafInterface

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainMenu struct {
	intractor *Interface

	name     *widget.Label
	start    *widget.Button
	load     *widget.Button
	settings *widget.Button
	exit     *widget.Button
}

func (c *MainMenu) Init(intractor *Interface) {
	c.intractor = intractor
	c.name = widget.NewLabel("        Checkers        ")
	c.start = widget.NewButton("Start", func() {
		c.intractor.Game.StartInit()
		c.intractor.Begin(&c.intractor.Game)
	})
	c.load = widget.NewButton("Load", func() {
		c.intractor.Begin(&c.intractor.Load)
	})
	c.settings = widget.NewButton("Settings", func() {})
	c.exit = widget.NewButton("Exit", intractor.Exit)
}

func (c *MainMenu) GetContent() fyne.CanvasObject {
	verticalBox2 := container.NewVBox(
		c.name,
		c.start,
		c.load,
		c.settings,
		c.exit,
	)
	content := container.New(
		layout.NewCenterLayout(),
		verticalBox2,
	)
	c.intractor.returnStatus = MainMenuStatus
	return content
}

func (c *MainMenu) KeyEventCallback(keyEvent *fyne.KeyEvent) {}
