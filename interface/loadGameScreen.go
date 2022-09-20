package _interface

import (
	"chekers/saveLoad"
	"fmt"
	"reflect"
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
	for save := range saveList {
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
	if command[len(command)-5:] == ".json" {
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
	fmt.Println("start resume", reflect.TypeOf(c))
	command := c.interactor.GetCommand(c.parse)
	c.interactor.switchCommander(command, c)
	fmt.Println("finish resume", reflect.TypeOf(c))
}
