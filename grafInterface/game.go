package grafInterface

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Game struct {
	intractor *Interface
}

func (c *Game) Init(intractor *Interface) {
	c.intractor = intractor
}

func (c *Game) GetContent() fyne.CanvasObject {
	c.intractor.returnStatus = MenuStatus
	return container.NewVBox()
}

func (c *Game) KeyEventCallback(keyEvent *fyne.KeyEvent) {
	if keyEvent.Name == fyne.KeyEscape {
		c.intractor.Return()
	}
}

func (c *Game) StartInit() {}
