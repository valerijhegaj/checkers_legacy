package _interface

import (
	"chekers/bot"
	"chekers/core"
	"chekers/gamer"
	"chekers/saveLoad"
	"fmt"
	"sync"
)

// interface status
const (
	mainMenu = iota
	menu
)

// commander name
const (
	returnToStatus = iota

	startGame
	loadGame
	game
	save
	exit

	display
	displayHelp
	resume
)

type Interface struct {
	gamer0 gamer.Gamer
	gamer1 gamer.Gamer
	bot0   bot.Bot
	bot1   bot.Bot

	mainMenuScreen
	startGameScreen
	loadGameScreen
	gameScreen
	menuScreen
	saveGameScreen

	status int

	exiter       chan int
	Participants saveLoad.Participants

	mutex sync.Mutex
}

func (c *Interface) switchCommander(command int, controler screenControler) {
	switch command {
	case returnToStatus:
		if c.status == mainMenu {
			go c.mainMenuScreen.Resume()
		} else {
			go c.menuScreen.Resume()
		}
	case startGame:
		go c.startGameScreen.Display()
	case loadGame:
		go c.loadGameScreen.Resume()
	case game:
		go c.gameScreen.Resume()
	case save:
		go c.saveGameScreen.Resume()
	case exit:
		go c.exit()
	case display:
		go controler.Display()
	case displayHelp:
		go controler.DisplayHelp()
	case resume:
		go controler.Resume()
	}
}

func (c *Interface) Init(exiter chan int, core core.GameCore) {
	c.exiter = exiter
	c.gamer0 = gamer.Gamer{0, &core}
	c.gamer1 = gamer.Gamer{1, &core}
	c.status = mainMenu
	c.bot0 = bot.CreateBot(0)
	c.bot1 = bot.CreateBot(0)

	c.mainMenuScreen.interactor = c
	c.startGameScreen.interactor = c
	c.loadGameScreen.interactor = c
	c.gameScreen.interactor = c
	c.menuScreen.interactor = c
	c.saveGameScreen.interactor = c

	c.switchCommander(display, c.mainMenuScreen)
}

func (c *Interface) GetCommand(parse func(string) int) int {
	var command string
	c.mutex.Lock()
	fmt.Scan(&command)
	c.mutex.Unlock()
	return parse(command)
}

func (c *Interface) initSave(save saveLoad.Save) {
	c.gamer0.InitSave(save)
	c.Participants = save.Master
	c.bot0 = bot.CreateBot(save.Master.Level0)
	c.bot1 = bot.CreateBot(save.Master.Level1)
}

func (c *Interface) CreateSave() saveLoad.Save {
	var save saveLoad.Save

	save.Field = c.gamer0.GetField()
	save.Master = c.Participants
	if c.gamer0.IsTurn() {
		save.TurnGamerId = 0
	} else {
		save.TurnGamerId = 1
	}

	return save
}

func (c *Interface) exit() {
	c.exiter <- 0
}
