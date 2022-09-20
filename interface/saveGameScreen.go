package _interface

import (
	"chekers/saveLoad"
	"fmt"
	"reflect"
)

type saveGameScreen struct {
	interactor *Interface
}

func (c saveGameScreen) Display() {
	fmt.Println("saves:")
	saveList, err := saveLoad.GetSaveList("saves")
	if err != nil {
		go c.Resume()
	}
	for save := range saveList {
		fmt.Println(save)
	}
	go c.Resume()
}

func (c saveGameScreen) DisplayHelp() {
	fmt.Println("save")
	displayHelpBasic()
	fmt.Println("path - print path where you want to save (game see saves only from saves)")
	go c.Resume()
}

func (c saveGameScreen) parse(command string) int {
	if command[len(command)-5:] == ".json" {
		err := c.writeSave(command)
		if err != nil {
			fmt.Println(err.Error())
			return resume
		}
		return returnToStatus
	}
	return parseBasic(command)
}

func (c saveGameScreen) Resume() {
	fmt.Println("start resume", reflect.TypeOf(c))
	command := c.interactor.GetCommand(c.parse)
	c.interactor.switchCommander(command, c)
	fmt.Println("finish resume", reflect.TypeOf(c))
}

func (c saveGameScreen) writeSave(path string) error {
	save := c.interactor.CreateSave()
	return save.Write(path)
}
