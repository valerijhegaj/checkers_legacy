package _interface

import (
	"chekers/core"
	"chekers/gamer"
	"chekers/saveLoad"
	"fmt"
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
	bot0   gamer.Bot
	bot1   gamer.Bot

	mainMenuScreen
	startGameScreen
	loadGameScreen
	gameScreen
	menuScreen
	saveGameScreen

	status int

	exiter chan int
	Master saveLoad.Master
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
	fmt.Println("start init interface")
	c.exiter = exiter
	c.gamer0 = gamer.Gamer{0, &core}
	c.gamer1 = gamer.Gamer{1, &core}
	c.status = mainMenu

	c.mainMenuScreen.interactor = c
	c.startGameScreen.interactor = c
	c.loadGameScreen.interactor = c
	c.gameScreen.interactor = c
	c.menuScreen.interactor = c
	c.saveGameScreen.interactor = c

	c.switchCommander(display, c.mainMenuScreen)
	fmt.Println("finish init interface")
}

func (c *Interface) GetCommand(parse func(string) int) int {
	var command string
	fmt.Scan(&command)
	return parse(command)
}

func (c *Interface) initSave(save saveLoad.Save) {
	c.gamer0.InitSave(save)
	c.Master = save.Master
}

func (c *Interface) CreateSave() saveLoad.Save {
	var save saveLoad.Save

	save.Field = c.gamer0.GetField()
	save.Master = c.Master
	if c.gamer0.IsTurn() {
		save.TurnGamerId = 0
	} else {
		save.TurnGamerId = 1
	}

	return save
}

func (c *Interface) exit() {
	fmt.Println("start exit interface")
	c.exiter <- 0
	fmt.Println("finish exit interface")
}
