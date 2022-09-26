package core

type checkersFeature struct {
	desk        *Field
	GamerEaters []int
}

func (c checkersFeature) Init() {
	c.GamerEaters = make([]int, 2)
}

func (c checkersFeature) CheckMove(from, to Coordinate, gamerId int) bool {
	if c.isGamerHasEater(gamerId) {
		return c.isMoveToEat(from, to)
	}
	return true
}

func (c *checkersFeature) MadeMove(from, to Coordinate, gamerId int) {
}

func (c checkersFeature) isCanBeEater(position Coordinate) bool {
	return c.isFoodInThisDirection(position, 1, 1) ||
		c.isFoodInThisDirection(position, 1, -1) ||
		c.isFoodInThisDirection(position, -1, -1) ||
		c.isFoodInThisDirection(position, -1, 1)
}

func (c checkersFeature) isFoodInThisDirection(from Coordinate, dx, dy int) bool {
	to := Coordinate{from.X + 2*dx, from.Y + 2*dy}
	figure := c.desk.At(from)
	if figure == nil {
		return false
	}
	for ; c.desk.InBorders(to); to.X, to.Y = to.X+dx, to.Y+dy {
		ok, foodPosition := c.isCanBeMove(from, to)
		if ok && foodPosition != c.desk.BordersLeft {
			return true
		}
	}
	return false
}

//start of optimization this feature
//func (c checkersFeature) isCanBeFood(position Coordinate) bool {
//	return c.isEaterInThisDirection(position, 1, 1) ||
//		c.isEaterInThisDirection(position, 1, -1) ||
//		c.isEaterInThisDirection(position, -1, -1) ||
//		c.isEaterInThisDirection(position, -1, 1)
//}
//
//
//func (c checkersFeature) isEaterInThisDirection(position Coordinate, dx, dy int) bool {
//	from := Coordinate{position.X + dx, position.Y + dy}
//	to := Coordinate{position.X - dx, position.Y - dy}
//
//	if !c.desk.IsAvailable(to) {
//		return false
//	}
//
//	for ; c.desk.InBorders(from); from.X, from.Y = from.X+dx, from.Y+dy {
//		ok, _ := c.isCanBeMove(from, to)
//		if ok {
//			return true
//		}
//	}
//	return false
//}

func (c checkersFeature) isMoveToEat(from, to Coordinate) bool {
	_, foodPosition := c.isCanBeMove(from, to)
	if foodPosition == c.desk.BordersLeft {
		return false
	}
	return true
}

func (c checkersFeature) isCanBeMove(from, to Coordinate) (bool, Coordinate) {
	figure := c.desk.At(from)
	if figure == nil {
		return false, c.desk.BordersLeft
	}
	return figure.IsMoveOne(c.desk, from, to)
}

func (c checkersFeature) isGamerHasEater(gamerId int) bool {
	for key, value := range c.desk.Figures {
		if value.GetOwnerId() == gamerId && c.isCanBeEater(key) {
			return true
		}
	}
	return false
}
