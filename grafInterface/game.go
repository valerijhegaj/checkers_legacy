package grafInterface

import (
	"chekers/core"
	"chekers/saveLoad"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"time"
)

type Game struct {
	intractor *Interface
}

func (c *Game) Init(intractor *Interface) {
	c.intractor = intractor
}

func (c *Game) GetContent() fyne.CanvasObject {
	c.intractor.returnStatus = MenuStatus
	return c.createBoard(c.intractor.gamer[0].GetField())
}

func (c *Game) KeyEventCallback(keyEvent *fyne.KeyEvent) {
	if keyEvent.Name == fyne.KeyEscape {
		c.intractor.Return()
	}
}

func (c *Game) StartInit() {
	var save saveLoad.Save
	save.Create()
	save.Master.Gamer0 = saveLoad.Man
	save.Master.Level0 = 4
	save.Master.Gamer1 = saveLoad.Bot
	save.Master.Level1 = 4

	c.intractor.InitSave(save)
}

func (c *Game) createBoard(board core.Field) *fyne.Container {
	boardDraw := container.NewGridWithColumns(
		board.BordersRight.Y - board.BordersLeft.Y + 1)
	var x, y int

	routine := func() {
		cellBg := canvas.NewRectangle(color.Gray{0x30})
		cell := NewCell(x, y, c.intractor)
		if x%2 == y%2 {
			cellBg.FillColor = color.Gray{0xE0}
		}
		figure := board.At(core.Coordinate{x, y})
		if figure != nil {
			img := canvas.NewImageFromResource(GetResource(figure))
			boardDraw.Add(container.NewMax(cellBg, cell, img))
		} else {
			boardDraw.Add(container.NewMax(cellBg, cell))
		}
	}
	if c.intractor.gamer[0].IsTurn() && c.intractor.Participants.Gamer0 == saveLoad.Man ||
		c.intractor.Participants.Gamer1 == saveLoad.Bot {
		for x = board.BordersRight.X; x >= board.BordersLeft.X; x-- {
			for y = board.BordersLeft.Y; y <= board.BordersRight.Y; y++ {
				routine()
			}
		}
	} else if c.intractor.gamer[1].IsTurn() && c.intractor.Participants.Gamer1 == saveLoad.Man ||
		c.intractor.Participants.Gamer1 == saveLoad.Bot {
		for x = board.BordersLeft.X; x <= board.BordersRight.X; x++ {
			for y = board.BordersRight.Y; y >= board.BordersLeft.Y; y-- {
				routine()
			}
		}
	}
	return boardDraw
}

func (c *Game) Routine() {
	if c.intractor.gamer[0].IsTurn() {
		if c.intractor.Participants.Gamer0 == saveLoad.Bot {
			c.intractor.bot[0].Move(c.intractor.gamer[0])
			defer c.intractor.Begin(&c.intractor.Game)
		}
	} else if c.intractor.Participants.Gamer1 == saveLoad.Bot {
		c.intractor.bot[1].Move(c.intractor.gamer[1])
		defer c.intractor.Begin(&c.intractor.Game)
	}
	time.Sleep(1000 * time.Microsecond)
}
