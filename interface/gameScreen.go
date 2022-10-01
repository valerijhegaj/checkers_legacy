package _interface

import (
	"fmt"
	"reflect"

	"chekers/bot"
	"chekers/core"
	"chekers/gamer"
	"chekers/saveLoad"
)

type gameScreen struct {
	interactor *Interface
}

func (c gameScreen) Display() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"

	field := c.interactor.gamer0.GetField()

	printFigureCorrect := func(x, y int) {
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

	if c.interactor.gamer0.IsTurn() {
		for x := field.BordersRight.X; x >= field.BordersLeft.X; x-- {
			fmt.Print(x+1, " ")
			for y := field.BordersLeft.Y; y <= field.BordersRight.Y; y++ {
				printFigureCorrect(x, y)
			}
			fmt.Println()
		}
	} else {
		for x := field.BordersLeft.X; x <= field.BordersRight.X; x++ {
			fmt.Print(8-x, " ")
			for y := field.BordersRight.Y; y >= field.BordersLeft.Y; y-- {
				printFigureCorrect(x, y)
			}
			fmt.Println()
		}
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
	fmt.Println(
		"move a4a5 - move figure from to, " +
			"for all gamers coordinates are absolute",
	)
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

func (c gameScreen) makeMove(
	gamer gamer.Gamer,
	from core.Coordinate,
	to []core.Coordinate,
) int {
	if gamer.Move(from, to) {
		return game
	} else {
		fmt.Println("incorrect move")
		return game
	}
}

func InitFromStringCoordinate(
	coordinate string,
	interactor *Interface,
) core.Coordinate {
	var c core.Coordinate
	c.X = int(coordinate[1] - '1')
	c.Y = int(coordinate[0] - 'a')
	if interactor.gamer1.IsTurn() {
		c.X = 7 - c.X
		c.Y = 7 - c.Y
	}
	return c
}

func ToStringCoordinate(
	c core.Coordinate,
	interactor *Interface,
) string {
	ans := ""
	if interactor.gamer1.IsTurn() {
		ans += string(rune(7 - c.Y + 'a'))
		ans += string(rune(7 - c.X + '1'))
	} else {

		ans += string(rune(c.Y + 'a'))
		ans += string(rune(c.X + '1'))
	}
	return ans
}

func (c gameScreen) getMove() (
	core.Coordinate,
	[]core.Coordinate,
) {
	var input string

	c.interactor.mutex.Lock()
	fmt.Scanln(&input)
	c.interactor.mutex.Unlock()

	if len(input)%2 != 0 {
		return core.Coordinate{0, 0}, []core.Coordinate{{0, 0}}
	}

	var coordinates []string
	for i := 0; i < len(input); i += 2 {
		coordinates = append(coordinates, input[i:i+2])
	}

	var from core.Coordinate
	var to []core.Coordinate

	for i, coordinate := range coordinates {
		if i == 0 {
			from = InitFromStringCoordinate(coordinate, c.interactor)
		} else {
			to = append(
				to,
				InitFromStringCoordinate(coordinate, c.interactor),
			)
		}
	}

	return from, to
}

func (c gameScreen) Resume() {
	if c.interactor.gamer0.IsTurn() {
		c.routine(
			c.interactor.Participants.Gamer0,
			c.interactor.gamer0,
			c.interactor.bot0,
		)
	} else {
		c.routine(
			c.interactor.Participants.Gamer1,
			c.interactor.gamer1,
			c.interactor.bot1,
		)
	}
}

func (c gameScreen) routine(
	master int,
	gamer gamer.Gamer,
	bot bot.Bot,
) {
	if master == saveLoad.Bot {
		from, to := bot.Move(gamer)
		c.print(from, to)
		c.interactor.switchCommander(game, c)
	} else {
		command := c.interactor.GetCommand(c.parse)
		c.interactor.switchCommander(command, c)
	}
}

func (c gameScreen) print(
	from core.Coordinate,
	to []core.Coordinate,
) {
	fmt.Print("from: ", ToStringCoordinate(from, c.interactor))
	fmt.Print(" to: ")
	for _, move := range to {
		fmt.Print(ToStringCoordinate(move, c.interactor), " ")
	}
	fmt.Println()
}
