package _interface

import (
	"fmt"
)

type mainMenuScreen struct {
	interactor *Interface
}

func (c mainMenuScreen) Display() {
	fmt.Println("checkers")
	fmt.Println("write 1 of 3 commands")
	fmt.Println("start")
	fmt.Println("load")
	fmt.Println("exit")
	go c.Resume()
}

func (c mainMenuScreen) DisplayHelp() {
	displayHelpBasicMenu()
	go c.Resume()
}

func (c mainMenuScreen) parse(command string) int {
	return parseBasicMenu(command)
}

func (c mainMenuScreen) Resume() {
	command := c.interactor.GetCommand(c.parse)
	c.interactor.switchCommander(command, c)
}
