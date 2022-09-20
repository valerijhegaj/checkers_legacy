package _interface

import (
	"fmt"
	"reflect"
)

type mainMenuScreen struct {
	interactor *Interface
}

func (c mainMenuScreen) Display() {
	fmt.Println("start display main menu")
	fmt.Println("checkers")
	fmt.Println("start")
	fmt.Println("load")
	fmt.Println("exit")
	go c.Resume()
	fmt.Println("finish display main menu")
}

func (c mainMenuScreen) DisplayHelp() {
	displayHelpBasicMenu()
	go c.Resume()
}

func (c mainMenuScreen) parse(command string) int {
	return parseBasicMenu(command)
}

func (c mainMenuScreen) Resume() {
	fmt.Println("start resume", reflect.TypeOf(c))
	command := c.interactor.GetCommand(c.parse)
	c.interactor.switchCommander(command, c)
	fmt.Println("finish resume", reflect.TypeOf(c))
}
