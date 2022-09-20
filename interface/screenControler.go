package _interface

import "fmt"

type screenControler interface {
	Display()
	DisplayHelp()
	Resume()
}

func parseBasic(command string) int {
	if command[0] == 'd' || command[1] == 'D' {
		return display
	} else if command[0] == 'r' || command[0] == 'R' {
		return returnToStatus
	}
	return displayHelp
}

func parseBasicMenu(command string) int {
	if command[0] == 's' || command[0] == 'S' {
		return startGame
	} else if command[0] == 'l' || command[0] == 'L' {
		return loadGame
	} else if command[0] == 'e' || command[0] == 'E' {
		return exit
	}
	return parseBasic(command)
}

func displayHelpBasic() {
	fmt.Println("help - get help")
	fmt.Println("display - display something about this")
	fmt.Println("return - return to menu")
}

func displayHelpBasicMenu() {
	fmt.Println("start - start new game")
	fmt.Println("load - load save")
	fmt.Println("exit - exit game")
	displayHelpBasic()
}
