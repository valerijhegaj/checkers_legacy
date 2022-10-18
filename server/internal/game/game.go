package game

import (
	"errors"

	"checkers/bot"
	"checkers/core"
	"checkers/gamer"
	"checkers/saveLoad"
	"checkers/server/internal/errorsStrings"
)

func NewGame(settings Settings, password string) *Game {
	var c core.GameCore
	game := Game{
		gamer: [2]gamer.Gamer{{0, &c}, {1, &c}},
		bot: [2]bot.Bot{
			bot.NewBot(settings.Level0), bot.NewBot(settings.Level1),
		},

		userID: [2]int{-1, -1},
		Participants: Participants{
			[2]int{settings.Gamer0, settings.Gamer1},
			[2]int{settings.Level0, settings.Level1},
		},

		password:   password,
		accessList: make(map[int]bool),

		winner: -1,
	}
	save := saveLoad.Save{
		Field:       core.NewStandard8x8Field(),
		TurnGamerId: 0,
	}
	game.gamer[0].InitSave(save)
	return &game
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

	winner int
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
		return errors.New(errorsStrings.IncorrectMove)
	}

	{
		isEnd, winner := c.gamer[0].GetWinner()
		if isEnd {
			c.winner = winner.GamerId
			return nil
		}
	}

	i = i ^ 1
	if c.Participants.gamer[i] == saveLoad.Bot {
		go func() {
			c.bot[i].Move(c.gamer[i])
			isEnd, winner := c.gamer[0].GetWinner()
			if isEnd {
				c.winner = winner.GamerId
			}
		}()
	}
	return nil
}

func (c *Game) GetGame(userID int) ([]byte, error) {
	var save saveLoad.Save
	if !c.accessList[userID] {
		return nil, errors.New(errorsStrings.PermissionDenied)
	}
	save.Field = c.gamer[0].GetField()
	if c.gamer[0].IsTurn() {
		save.TurnGamerId = 0
	}
	save.TurnGamerId = 1

	save.Master.Gamer0 = c.Participants.gamer[0]
	save.Master.Gamer1 = c.Participants.gamer[1]
	save.Master.Level0 = c.Participants.level[0]
	save.Master.Level1 = c.Participants.level[1]

	save.Winner = c.winner
	return save.GetRawSave()
}

func (c *Game) AddUser(userID int, password string) error {
	if c.password != password {
		return errors.New(errorsStrings.PermissionDenied)
	}
	defer func() { c.accessList[userID] = true }()
	if c.Participants.gamer[0] == saveLoad.Bot {
		if c.userID[1] == -1 {
			go c.bot[0].Move(c.gamer[0])
			c.userID[1] = userID
		}
		return nil
	}
	if c.userID[0] == -1 {
		c.userID[0] = userID
		return nil
	}
	if c.Participants.gamer[1] == saveLoad.Man && c.userID[1] == -1 {
		c.userID[1] = userID
		return nil
	}
	return nil
}
