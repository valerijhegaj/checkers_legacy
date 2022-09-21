package _interface

import (
	"fmt"
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
	go c.Resume()
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
	command := c.interactor.GetCommand(c.parse)
	c.interactor.switchCommander(command, c)
}
