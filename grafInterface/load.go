package grafInterface

import (
	"fmt"

	"checkers/saveLoad"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Load struct {
	intractor *Interface
}

func (c *Load) Init(intractor *Interface) {
	c.intractor = intractor
}

func (c *Load) GetContent() fyne.CanvasObject {
	saveList, err := saveLoad.GetSaveList("saves")
	if err != nil || len(saveList) == 0 {
		label := widget.NewLabel("not found saves :(")
		if err != nil {
			fmt.Println(err.Error())
		}
		return container.NewVBox(label)
	}
	content := container.NewVBox()
	for i := range saveList {
		name := saveList[i]
		content.Add(
			widget.NewButton(
				name[:len(name)-5], func() {
					c.Load(name)
				},
			),
		)
	}
	return container.NewVScroll(
		content,
	)
}

func (c *Load) KeyEventCallback(keyEvent *fyne.KeyEvent) {
	if keyEvent.Name == fyne.KeyEscape {
		c.intractor.Return()
	}
}

func (c *Load) Load(name string) {
	path := "saves/" + name
	var save saveLoad.Save
	err := save.Read(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.intractor.InitSave(save)
	c.intractor.Begin(&c.intractor.Game)
}
