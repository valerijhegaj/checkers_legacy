package core

type Checker struct {
	OwnerId int
}

func (c Checker) GetOwnerId() int {
	return c.OwnerId
}

func (c Checker) newTransKing(desk *Field, position Coordinate) King {
	desk.RemoveWithOutBin(position)
	king := King{c.OwnerId}
	desk.Put(position, king)
	return king
}

func (c Checker) getVerticalToTransKing(desk *Field) int {
	vertical := desk.BordersRight.X
	if c.OwnerId == 1 {
		vertical = desk.BordersLeft.X
	}
	return vertical
}

func (c Checker) Move(
	desk *Field,
	from Coordinate,
	way []Coordinate,
) (
	bool,
	Coordinate,
) {
	vertical := c.getVerticalToTransKing(desk)

	isMoveWithoutEat, _ := c.isMoveWithoutEat(desk, from, way[0])
	if isMoveWithoutEat {
		desk.Move(from, way[0])

		if way[0].X == vertical {
			c.newTransKing(desk, way[0])
		}
		return true, way[0]
	}

	isCanBeMoved, foodPosition := c.isMoveToEat(desk, from, way[0])
	if !isCanBeMoved {
		return false, from
	}
	desk.Remove(foodPosition)
	desk.Move(from, way[0])

	if way[0].X == vertical {
		king := c.newTransKing(desk, way[0])
		return king.moveToEat(desk, way[0], way[1:])
	}
	return c.moveToEat(desk, way[0], way[1:])
}

func (c Checker) moveToEat(
	desk *Field,
	from Coordinate,
	way []Coordinate,
) (
	bool,
	Coordinate,
) {
	vertical := c.getVerticalToTransKing(desk)
	for i, to := range way {
		isCanBeMoved, foodPosition := c.isMoveToEat(desk, from, to)

		if !isCanBeMoved {
			return true, from
		}

		desk.Remove(foodPosition)
		desk.Move(from, to)

		from = to
		if to.X == vertical {
			king := c.newTransKing(desk, to)
			return king.moveToEat(desk, to, way[i+1:])
		}
	}
	return true, from
}

func (c Checker) IsMoveOne(desk *Field, from, to Coordinate) (
	bool,
	Coordinate,
) {
	isMoveWithFood, foodPosition := c.isMoveToEat(desk, from, to)
	if isMoveWithFood {
		return isMoveWithFood, foodPosition
	}
	isMoveWithOutFood, _ := c.isMoveWithoutEat(desk, from, to)
	return isMoveWithOutFood, foodPosition
}

func (c Checker) isMoveWithoutEat(
	desk *Field,
	from, to Coordinate,
) (
	bool,
	Coordinate,
) {
	vertical := 1
	if c.GetOwnerId() == 1 {
		vertical = -1
	}
	if to.X-from.X == vertical &&
		(to.Y-from.Y == 1 || to.Y-from.Y == -1) {
		if desk.IsAvailable(to) {
			return true, desk.BordersLeft
		}
	}
	return false, desk.BordersLeft
}

func (c Checker) isMoveToEat(desk *Field, from, to Coordinate) (
	bool,
	Coordinate,
) {
	foodPosition := Coordinate{
		(to.X + from.X) / 2,
		(to.Y + from.Y) / 2,
	}

	if (to.X-from.X == 2 || to.X-from.X == -2) &&
		(to.Y-from.Y == 2 || to.Y-from.Y == -2) {
		if desk.IsAvailable(to) && !desk.IsAvailable(foodPosition) {
			food := desk.At(foodPosition)
			if food.GetOwnerId() != c.GetOwnerId() {
				return true, foodPosition
			}
		}
	}

	return false, desk.BordersLeft
}

func (c Checker) addMove(
	moves *[]Coordinate,
	desk *Field,
	d, from Coordinate,
	IsMove func(field *Field, from, to Coordinate) (
		bool,
		Coordinate,
	),
) {
	move := Coordinate{from.X + d.X, from.Y + d.Y}
	isMove, _ := IsMove(desk, from, move)
	if isMove {
		*moves = append(*moves, move)
	}
}

func (c Checker) GetAvailableMoves(
	desk *Field,
	from Coordinate,
) []Coordinate {
	var moves []Coordinate
	c.addMove(&moves, desk, Coordinate{1, 1}, from, c.isMoveWithoutEat)
	c.addMove(&moves, desk, Coordinate{1, -1}, from, c.isMoveWithoutEat)
	c.addMove(&moves, desk, Coordinate{-1, 1}, from, c.isMoveWithoutEat)
	c.addMove(
		&moves,
		desk,
		Coordinate{-1, -1},
		from,
		c.isMoveWithoutEat,
	)
	c.addMove(&moves, desk, Coordinate{2, 2}, from, c.isMoveToEat)
	c.addMove(&moves, desk, Coordinate{-2, 2}, from, c.isMoveToEat)
	c.addMove(&moves, desk, Coordinate{2, -2}, from, c.isMoveToEat)
	c.addMove(&moves, desk, Coordinate{-2, -2}, from, c.isMoveToEat)
	return moves
}

func (c Checker) GetAvailableMovesToEat(
	desk *Field,
	from Coordinate,
) []Coordinate {
	var moves []Coordinate
	c.addMove(&moves, desk, Coordinate{2, 2}, from, c.isMoveToEat)
	c.addMove(&moves, desk, Coordinate{-2, 2}, from, c.isMoveToEat)
	c.addMove(&moves, desk, Coordinate{2, -2}, from, c.isMoveToEat)
	c.addMove(&moves, desk, Coordinate{-2, -2}, from, c.isMoveToEat)
	return moves
}
