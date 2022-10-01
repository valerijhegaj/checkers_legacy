package _interface

import (
	"fmt"

	"chekers/saveLoad"
)

type loadGameScreen struct {
	interactor *Interface
}

func (c loadGameScreen) Display() {
	saveList, err := saveLoad.GetSaveList("saves")
	if err != nil {
		fmt.Println(err.Error())
		go c.Resume()
	}
	fmt.Println("saves:")
	for _, save := range saveList {
		fmt.Println(save)
	}
	go c.Resume()
}

func (c loadGameScreen) DisplayHelp() {
	fmt.Println("path - path to json file which is save")
	displayHelpBasic()
	go c.Resume()
}

func (c loadGameScreen) parse(command string) int {
	if len(command) >= 5 && command[len(command)-5:] == ".json" {
		var save saveLoad.Save
		err := save.Read(command)
		if err != nil {
			fmt.Println(err.Error())
			return resume
		}

		c.interactor.initSave(save)

		c.interactor.status = menu
		return game
	}
	return parseBasic(command)
}

func (c loadGameScreen) Resume() {
	command := c.interactor.GetCommand(c.parse)
	c.interactor.switchCommander(command, c)
}
