package game

import (
	"errors"

	"checkers/bot"
	"checkers/core"
	"checkers/gamer"
	"checkers/saveLoad"
	"checkers/server/internal/errorsStrings"
)

func NewGame(settings Settings, password string) Game {
	var c core.GameCore
	return Game{
		gamer: [2]gamer.Gamer{{0, &c}, {1, &c}},
		bot: [2]bot.Bot{
			bot.NewBot(settings.Level0), bot.NewBot(settings.Level1),
		},
		Participants: Participants{
			[2]int{settings.Gamer0, settings.Gamer1},
			[2]int{settings.Level0, settings.Level1},
		},
		password:   password,
		accessList: make(map[int]bool),
	}
}

type Participants struct {
	gamer [2]int
	level [2]int
}

type Settings saveLoad.Participants

type Game struct {
	gamer [2]gamer.Gamer
	bot   [2]bot.Bot

	userID [2]int
	Participants

	password   string
	accessList map[int]bool
}

func (c *Game) Move(
	userID int, from core.Coordinate, path []core.Coordinate,
) error {
	var i int
	switch userID {
	case c.userID[0]:
		i = 0
	case c.userID[1]:
		i = 1
	default:
		return errors.New(errorsStrings.PermissionDenied)
	}
	if !c.gamer[i].Move(from, path) {
		return errors.New(errorsStrings.PermissionDenied)
	}
	i = i ^ 1
	if c.Participants.gamer[i] == saveLoad.Bot {
		go c.bot[i].Move(c.gamer[i])
	}
	return nil
}

func (c *Game) GetGame(userID int) (saveLoad.Save, error) {
	return saveLoad.Save{}, nil
}
