package _interface

import (
	"chekers/core"
	"chekers/gamer"
	"chekers/saveLoad"
	"fmt"
	"reflect"
)

type gameScreen struct {
	interactor *Interface
}

func (c gameScreen) Display() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"

	field := c.interactor.gamer0.GetField()
	for x := field.BordersRight.X; x >= field.BordersLeft.X; x-- {
		fmt.Print(x+1, " ")
		for y := field.BordersLeft.Y; y <= field.BordersRight.Y; y++ {
			figure := field.At(core.Coordinate{x, y})
			if figure == nil {
				fmt.Print("_ ")
			} else if reflect.TypeOf(figure) == reflect.TypeOf(core.Checker{}) {
				if figure.GetOwnerId() == 1 {
					fmt.Print(colorRed, "O ", colorReset)
				} else {
					fmt.Print("O ")
				}
			} else {
				if figure.GetOwnerId() == 1 {
					fmt.Print(colorRed, "K ", colorReset)
				} else {
					fmt.Print("K ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Print("  ")
	for y := field.BordersLeft.Y; y <= field.BordersRight.Y; y++ {
		fmt.Print(string(rune('a'+y)), " ")
	}
	fmt.Println()
	go c.Resume()
}

func (c gameScreen) DisplayHelp() {
	displayHelpBasic()
	fmt.Println("move a4a5 - move figure from to, " +
		"for all gamers coordinates are absolute")
	fmt.Println("O - checker, K - king, _ - empty")
	go c.Resume()
}

func (c gameScreen) parse(command string) int {
	if command == "move" || command == "Move" {
		from, to := c.getMove()
		if c.interactor.gamer0.IsTurn() {
			return c.makeMove(c.interactor.gamer0, from, to)
		} else {
			return c.makeMove(c.interactor.gamer1, from, to)
		}
	}
	return parseBasic(command)
}

func (c gameScreen) makeMove(gamer gamer.Gamer, from core.Coordinate, to []core.Coordinate) int {
	if gamer.Move(from, to) {
		return game
	} else {
		fmt.Println("incorrect move")
		return game
	}
}

func (c gameScreen) getMove() (core.Coordinate, []core.Coordinate) {
	var input string
	c.interactor.mutex.Lock()
	fmt.Scanln(&input)
	c.interactor.mutex.Unlock()

	var coordinates []string
	for i := 0; i < len(input); i += 2 {
		coordinates = append(coordinates, input[i:i+2])
	}

	var from core.Coordinate
	var to []core.Coordinate

	from.InitFromString(coordinates[0])
	for i, coordinate := range coordinates {
		if i == 0 {
			from.InitFromString(coordinate)
		} else {
			var loacalTo core.Coordinate
			loacalTo.InitFromString(coordinate)
			to = append(to, loacalTo)
		}
	}

	return from, to
}

func (c gameScreen) Resume() {
	if c.interactor.gamer0.IsTurn() {
		c.routine(c.interactor.Master.Gamer0, c.interactor.gamer0, c.interactor.bot0)
	} else {
		c.routine(c.interactor.Master.Gamer1, c.interactor.gamer1, c.interactor.bot1)
	}
}

func (c gameScreen) routine(master int, gamer gamer.Gamer, bot gamer.Bot) {
	if master == saveLoad.Bot {
		bot.Move(gamer)
	} else {
		command := c.interactor.GetCommand(c.parse)
		c.interactor.switchCommander(command, c)
	}
}
