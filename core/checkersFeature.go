package core

type checkersFeature struct{}

func (c checkersFeature) IsCanBeEater(desk *Field, position Coordinate) bool {
	return c.isFoodInThisDirection(desk, position, 1, 1) ||
		c.isFoodInThisDirection(desk, position, 1, -1) ||
		c.isFoodInThisDirection(desk, position, -1, -1) ||
		c.isFoodInThisDirection(desk, position, -1, 1)
}

func (c checkersFeature) IsCanBeFood(desk *Field, position Coordinate) bool {
	return c.isEaterInThisDirecytion(desk, position, 1, 1) ||
		c.isEaterInThisDirecytion(desk, position, 1, -1) ||
		c.isEaterInThisDirecytion(desk, position, -1, -1) ||
		c.isEaterInThisDirecytion(desk, position, -1, 1)
}

func (c checkersFeature) isFoodInThisDirection(desk *Field, from Coordinate, dx, dy int) bool {
	to := Coordinate{from.X + 2*dx, from.Y + 2*dy}

	for desk.InBorders(to) {
		break
	}
	//not implemented
	return true
}

func (c checkersFeature) isEaterInThisDirecytion(desk *Field, position Coordinate, dx, dy int) bool {
	from := Coordinate{position.X + dx, position.Y + dy}
	to := Coordinate{position.X - dx, position.Y - dy}

	if !desk.IsFree(to) {
		return false
	}

	for desk.InBorders(from) {
		if c.isCanBeMove(desk, from, to) {
			return true
		}

		from.X += dx
		from.Y += dy
	}
	return false
}

func (c checkersFeature) isCanBeMove(desk *Field, from, to Coordinate) bool {
	figure := desk.At(from)
	if figure == nil {
		return false
	}
	return figure.IsMoveOne(desk, from, to)
}
