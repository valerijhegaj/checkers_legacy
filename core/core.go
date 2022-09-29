package core

import "sync"

type GameCore struct {
	field       Field
	turnGamerId int
	checkersFeature
	Mutex sync.Mutex
}

func (c GameCore) GetField() Field {
	return c.field
}

func (c GameCore) IsTurn(gamerId int) bool {
	return gamerId == c.turnGamerId
}

func (c *GameCore) Move(from Coordinate, way []Coordinate, gamerId int) bool {
	c.Mutex.Lock()
	if gamerId != c.turnGamerId {
		return false
	}
	figure := c.field.At(from)
	if figure == nil {
		return false
	}
	if figure.GetOwnerId() != c.turnGamerId {
		return false
	}
	if !c.checkersFeature.CheckMove(from, way[0], gamerId) {
		return false
	}

	sucсeess, to := figure.Move(&c.field, from, way)
	if sucсeess {
		c.checkersFeature.MadeMove(from, to, gamerId)
		c.turnGamerId ^= 1
	}
	c.Mutex.Unlock()
	return sucсeess
}

func (c *GameCore) InitField(field Field) {
	c.field = field
	c.checkersFeature.desk = &field
}

func (c *GameCore) InitTurnGamerId(turnGamerId int) {
	c.turnGamerId = turnGamerId
}
