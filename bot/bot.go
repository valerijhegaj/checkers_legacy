package bot

import (
	"chekers/core"
	"chekers/gamer"
)

type Bot struct {
	Analyzator
}

func (c *Bot) Move(gamer gamer.Gamer) {
	field := gamer.GetField()
	from, way := c.analyzeField(&field, gamer.GamerId)
	gamer.Move(from, way)
}

type Analyzator interface {
	analyzeField(field *core.Field, gamerId int) (core.Coordinate, []core.Coordinate)
}
