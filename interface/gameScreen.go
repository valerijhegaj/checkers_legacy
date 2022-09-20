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
	field := c.interactor.gamer0.GetField()
	for x := 7; x >= 0; x-- {
		for y := 0; y < 8; y++ {
			figure := field.At(core.Coordinate{x, y})
			if figure == nil {
				fmt.Print("0 ")
			} else if reflect.TypeOf(figure) == reflect.TypeOf(core.Checker{}) {
				fmt.Print("1 ")
			} else {
				fmt.Print("2 ")
			}
		}
		fmt.Println()
	}
	go c.Resume()
}

func (c gameScreen) DisplayHelp() {
	displayHelpBasic()
	fmt.Println("move a4 a5 - move figure from to, " +
		"for all gamers coordinates are absolute, " +
		"from view of white left bottom is a0, " +
		"chars in horizontal, nums in vertical")

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

func (c gameScreen) makeMove(gamer gamer.Gamer, from, to core.Coordinate) int {
	if gamer.Move(from, to) {
		return game
	} else {
		fmt.Println("incorrect move")
		return game
	}
}

func (c gameScreen) getMove() (core.Coordinate, core.Coordinate) {
	var input string
	fmt.Scan(&input)
	var from, to core.Coordinate
	from.InitFromString(input)
	fmt.Scan(&input)
	to.InitFromString(input)

	return from, to
}

func (c gameScreen) Resume() {
	fmt.Println("start resume", reflect.TypeOf(c))
	if c.interactor.gamer0.IsTurn() {
		c.routine(c.interactor.Master.Gamer0, c.interactor.gamer0, c.interactor.bot0)
	} else {
		c.routine(c.interactor.Master.Gamer1, c.interactor.gamer1, c.interactor.bot1)
	}
	fmt.Println("finish resume", reflect.TypeOf(c))
}

func (c gameScreen) routine(master int, gamer gamer.Gamer, bot gamer.Bot) {
	if master == saveLoad.Bot {
		bot.Move(gamer)
	} else {
		command := c.interactor.GetCommand(c.parse)
		c.interactor.switchCommander(command, c)
	}
}
