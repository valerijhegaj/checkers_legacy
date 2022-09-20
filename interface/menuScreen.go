package _interface

import (
	"fmt"
	"reflect"
)

type menuScreen struct {
	interactor *Interface
}

func (c menuScreen) Display() {
	fmt.Println("resume")
	fmt.Println("start")
	fmt.Println("save")
	fmt.Println("load")
	fmt.Println("exit")
	go c.Resume()
}

func (c menuScreen) DisplayHelp() {
	fmt.Println("resume - return to game")
	fmt.Println("save - save game")
	displayHelpBasicMenu()
}

func (c menuScreen) parse(command string) int {
	if command == "resume" || command == "Resume" {
		return game
	} else if command == "save" || command == "Save" {
		return save
	}
	return parseBasicMenu(command)
}

func (c menuScreen) Resume() {
	fmt.Println("start resume", reflect.TypeOf(c))
	command := c.interactor.GetCommand(c.parse)
	c.interactor.switchCommander(command, c)
	fmt.Println("finish resume", reflect.TypeOf(c))
}
