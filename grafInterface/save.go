package grafInterface

import (
	"fmt"

	"checkers/saveLoad"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Save struct {
	intractor *Interface

	save *widget.Button

	name *string
}

func (c *Save) Init(intractor *Interface) {
	c.intractor = intractor

	c.save = widget.NewButton(
		"Save", func() {
			c.Save(c.name)
		},
	)
}

func (c *Save) GetContent() fyne.CanvasObject {
	saveList, err := saveLoad.GetSaveList("saves")
	content := container.NewVBox()

	input := widget.NewEntry()
	input.SetPlaceHolder("Save name")
	content.Add(input)
	c.name = &input.Text
	content.Add(c.save)
	if err != nil || len(saveList) == 0 {
		content.Add(widget.NewLabel("not found saves :("))
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		for _, label := range saveList {
			content.Add(widget.NewLabel(label[:len(label)-5]))
		}
	}
	return container.NewVScroll(
		content,
	)
}

func (c *Save) KeyEventCallback(keyEvent *fyne.KeyEvent) {
	if keyEvent.Name == fyne.KeyEscape {
		c.intractor.Return()
	} else if keyEvent.Name == fyne.KeyEnter {
		c.save.Tapped(&fyne.PointEvent{})
	}
}

func (c *Save) Save(name *string) {
	path := "saves/" + *name + ".json"
	save := c.intractor.CreateSave()
	err := save.Write(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.intractor.Return()
}
