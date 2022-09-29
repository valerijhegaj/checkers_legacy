package bot

import (
	"chekers/core"
	"chekers/gamer"
)

func NewBot(level int) Bot {
	var bot Bot
	if level == 0 {
		bot.Analyzator = CreateRandomMoves()
	} else {
		bot.Analyzator = minMax{level, tree{}}
	}
	return bot
}

type Bot struct {
	Analyzator
}

func (c *Bot) Move(gamer gamer.Gamer) (core.Coordinate, []core.Coordinate) {
	field := gamer.GetField()
	from, way := c.analyzeField(&field, gamer.GamerId)

	gamer.Move(from, way)
	return from, way
}

type Analyzator interface {
	analyzeField(field *core.Field, gamerId int) (core.Coordinate, []core.Coordinate)
}
