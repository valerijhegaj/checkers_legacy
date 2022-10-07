package data

import (
	"checkers/bot"
	"checkers/core"
	"checkers/gamer"
	"checkers/saveLoad"
)

func NewGame(
	save saveLoad.Save, password string, gamerID int,
	user string,
) Game {
	var c Game
	c.gamer[0] = gamer.Gamer{0, &c.c}
	c.gamer[1] = gamer.Gamer{1, &c.c}
	c.bot[0] = bot.NewBot(save.Master.Level0)
	c.bot[1] = bot.NewBot(save.Master.Level1)
	c.Participants = save.Master
	c.user[gamerID] = user
	c.password = password

	c.gamer[0].InitSave(save)
	return c
}

type Game struct {
	c        core.GameCore
	gamer    [2]gamer.Gamer
	bot      [2]bot.Bot
	user     [2]string
	password string

	saveLoad.Participants
}

func (c *Game) Init(save saveLoad.Save) {

}

func (c *Game) Move(
	from core.Coordinate, way []core.Coordinate,
	gamerId int,
) bool {
	return c.gamer[gamerId].Move(from, way)
}

func (c *Game) GetField() core.Field {
	return c.gamer[0].GetField()
}

func (c *Game) GetSave() saveLoad.Save {
	var save saveLoad.Save
	save.Field = c.gamer[0].GetField()
	save.Master = c.Participants
	if c.gamer[0].IsTurn() {
		save.TurnGamerId = 0
	} else {
		save.TurnGamerId = 1
	}
	return save
}
