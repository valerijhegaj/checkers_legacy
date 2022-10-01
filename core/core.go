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

func (c *GameCore) Move(
	from Coordinate,
	way []Coordinate,
	gamerId int,
) bool {
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

	success, _ := figure.Move(&c.field, from, way)
	if success {
		c.turnGamerId ^= 1
	}
	return success
}

func (c *GameCore) InitField(field Field) {
	c.field = field
	c.checkersFeature.desk = &field
}

func (c *GameCore) InitTurnGamerId(turnGamerId int) {
	c.turnGamerId = turnGamerId
}
