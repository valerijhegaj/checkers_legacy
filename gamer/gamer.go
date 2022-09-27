package gamer

import (
	"chekers/core"
	"chekers/saveLoad"
)

type Gamer struct {
	GamerId int
	Core    *core.GameCore
}

func (c Gamer) GetField() core.Field {
	return c.Core.GetField()
}

func (c Gamer) IsTurn() bool {
	return c.Core.IsTurn(c.GamerId)
}

func (c Gamer) Move(from core.Coordinate, way []core.Coordinate) bool {
	return c.Core.Move(from, way, c.GamerId)
}

func (c Gamer) InitSave(save saveLoad.Save) {
	c.Core.InitField(save.Field)
	c.Core.InitTurnGamerId(save.TurnGamerId)
}
