package grafInterface

import (
	"chekers/bot"
	"chekers/core"
	"chekers/gamer"
	"chekers/saveLoad"
	"fyne.io/fyne/v2"
)

const (
	MainMenuStatus = iota
	MenuStatus
)

type Interface struct {
	a *fyne.App
	w *fyne.Window

	MainMenu
	Menu
	Load
	Save
	Game

	gamer [2]gamer.Gamer

	bot [2]bot.Bot

	Participants saveLoad.Participants
	returnStatus int
}

func (c *Interface) Init(a *fyne.App, w *fyne.Window, core *core.GameCore) {
	c.a = a
	c.w = w

	c.MainMenu.Init(c)
	c.Menu.Init(c)
	c.Load.Init(c)
	c.Save.Init(c)
	c.Game.Init(c)

	c.gamer[0] = gamer.Gamer{0, core}
	c.gamer[1] = gamer.Gamer{1, core}
	c.bot[0] = bot.NewBot(3)
	c.bot[1] = bot.NewBot(3)

	c.returnStatus = MainMenuStatus
}

type Commander interface {
	GetContent() fyne.CanvasObject
	KeyEventCallback(keyEvent *fyne.KeyEvent)
}

func (c *Interface) Begin(commander Commander) {
	content := commander.GetContent()
	(*c.w).SetContent(content)
	(*c.w).Canvas().SetOnTypedKey(commander.KeyEventCallback)
}

func (c Interface) Exit() {
	(*c.a).Quit()
}

func (c *Interface) InitSave(save saveLoad.Save) {
	c.gamer[0].InitSave(save)
	c.Participants = save.Master
	c.bot[0] = bot.NewBot(save.Master.Level0)
	c.bot[1] = bot.NewBot(save.Master.Level1)
}

func (c *Interface) CreateSave() saveLoad.Save {
	var save saveLoad.Save

	save.Field = c.gamer[0].GetField()
	save.Master = c.Participants
	if c.gamer[0].IsTurn() {
		save.TurnGamerId = 0
	} else {
		save.TurnGamerId = 1
	}

	return save
}

func (c *Interface) Return() {
	if c.returnStatus == MainMenuStatus {
		c.Begin(&c.MainMenu)
	} else {
		c.Begin(&c.Menu)
	}
}
