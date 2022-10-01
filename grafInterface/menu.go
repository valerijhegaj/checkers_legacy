package grafInterface

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Menu struct {
	intractor *Interface

	name     *widget.Label
	resume   *widget.Button
	start    *widget.Button
	save     *widget.Button
	load     *widget.Button
	settings *widget.Button
	exit     *widget.Button
}

func (c *Menu) Init(intractor *Interface) {
	c.intractor = intractor

	c.name = intractor.MainMenu.name
	c.resume = widget.NewButton(
		"Resume", func() {
			c.intractor.Begin(&intractor.Game)
		},
	)
	c.start = intractor.MainMenu.start
	c.save = widget.NewButton(
		"Save", func() {
			c.intractor.Begin(&intractor.Save)
		},
	)
	c.load = intractor.MainMenu.load
	c.settings = intractor.MainMenu.settings
	c.exit = intractor.MainMenu.exit
}

func (c *Menu) GetContent() fyne.CanvasObject {
	verticalBox2 := container.NewVBox(
		c.name,
		c.resume,
		c.start,
		c.save,
		c.load,
		c.settings,
		c.exit,
	)
	content := container.New(
		layout.NewCenterLayout(),
		verticalBox2,
	)
	c.intractor.returnStatus = MenuStatus
	return content
}

func (c *Menu) KeyEventCallback(keyEvent *fyne.KeyEvent) {
	if keyEvent.Name == fyne.KeyEscape {
		c.resume.Tapped(&fyne.PointEvent{})
	}
}
