package core

type checkersFeature struct {
	desk *Field
}

func (c checkersFeature) CheckMove(
	from, to Coordinate,
	gamerId int,
) bool {
	if c.isGamerHasEater(gamerId) {
		return c.isMoveToEat(from, to)
	}
	return true
}

func (c checkersFeature) isGamerHasEater(gamerId int) bool {
	for key, figure := range c.desk.Figures {
		if figure.GetOwnerId() == gamerId {
			if figure.GetAvailableMovesToEat(c.desk, key) != nil {
				return true
			}
		}
	}
	return false
}

func (c checkersFeature) isMoveToEat(from, to Coordinate) bool {
	_, foodPosition := c.desk.At(from).IsMoveOne(c.desk, from, to)
	if foodPosition == c.desk.BordersLeft {
		return false
	}
	return true
}
