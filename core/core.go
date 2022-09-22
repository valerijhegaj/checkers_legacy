package core

type GameCore struct {
	field       Field
	turnGamerId int
	checkersFeature
}

func (c GameCore) GetField() Field {
	return c.field
}

func (c GameCore) IsTurn(gamerId int) bool {
	return gamerId == c.turnGamerId
}

func (c *GameCore) Move(from Coordinate, to []Coordinate, gamerId int) bool {
	figure := c.field.At(from)
	if figure == nil {
		return false
	}
	if figure.GetOwnerId() != gamerId {
		return false
	}
	sucsees, _ := figure.Move(&c.field, from, to)
	if sucsees {
		c.turnGamerId += 1
		c.turnGamerId %= 2
	}
	return sucsees
}

func (c *GameCore) InitField(field Field) {
	c.field = field
}

func (c *GameCore) InitTurnGamerId(turnGamerId int) {
	c.turnGamerId = turnGamerId
}
