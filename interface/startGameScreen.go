package _interface

import (
	"chekers/saveLoad"
	"fmt"
	"reflect"
)

type startGameScreen struct {
	interactor *Interface
}

func (c startGameScreen) Display() {
	fmt.Println("print start")
	fmt.Println("print gamer or bot (who will be for white)")
	fmt.Println("print gamer or bot (who will be for nigga)")
	go c.Resume()
}

func (c startGameScreen) DisplayHelp() {
	displayHelpBasic()
	go c.Display()
}

func (c startGameScreen) parse(command string) int {

	if command == "start" || command == "Start" {
		var save saveLoad.Save
		err := save.Read("startFields/start_field.json")
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("can't open start field, make shure you install all right and didn't delete anything")
			return resume
		}

		save.Master = c.getMaster()
		save.TurnGamerId = 0
		c.interactor.initSave(save)

		c.interactor.status = menu

		return game
	}
	return parseBasic(command)
}

func (c startGameScreen) getMaster() saveLoad.Master {
	var master saveLoad.Master
	var name string
	fmt.Scan(&name)
	c.parseMasterOne(&master.Gamer0, name)
	fmt.Scan(&name)
	c.parseMasterOne(&master.Gamer1, name)
	return master
}

func (c startGameScreen) parseMasterOne(gamer *int, name string) {
	if name == "bot" || name == "Bot" {
		*gamer = saveLoad.Bot
	} else {
		*gamer = saveLoad.Man
	}
}

func (c startGameScreen) Resume() {
	fmt.Println("start resume", reflect.TypeOf(c))
	command := c.interactor.GetCommand(c.parse)
	c.interactor.switchCommander(command, c)
	fmt.Println("finish resume", reflect.TypeOf(c))
}
