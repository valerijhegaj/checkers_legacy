package bot

import (
	"chekers/core"
	"chekers/gamer"
)

func NewBot(level int) Bot {
	var bot Bot
	if level == 0 {
		bot.Mind = NewRandomMoves()
	} else if level <= 4 {
		bot.Mind = NewMinMax(level)
	} else {
		bot.Mind = NewMinMaxV2(level-1, 3, 1)
	}
	return bot
}

type Bot struct {
	Mind
}

func (c *Bot) Move(gamer gamer.Gamer) (
	core.Coordinate,
	[]core.Coordinate,
) {
	field := gamer.GetField()
	from, way := c.GetMove(&field, gamer.GamerId)

	gamer.Move(from, way)
	return from, way
}

type Mind interface {
	GetMove(field *core.Field, gamerId int) (
		core.Coordinate,
		[]core.Coordinate,
	)
}
