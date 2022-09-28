package core

type Checker struct {
	OwnerId int
}

func (c Checker) GetAvailableMoves(desk *Field, from Coordinate) []Coordinate {
	var moves []Coordinate
	addMove := func(dx, dy int, checker func(field *Field, from, to Coordinate) (bool, Coordinate)) {
		move := Coordinate{from.X + dx, from.Y + dy}
		isMove, _ := checker(desk, from, move)
		if isMove {
			moves = append(moves, move)
		}
	}

	addMove(1, 1, c.isMoveWithoutEat)
	addMove(-1, 1, c.isMoveWithoutEat)
	addMove(1, -1, c.isMoveWithoutEat)
	addMove(-1, -1, c.isMoveWithoutEat)
	addMove(2, 2, c.isMoveToEat)
	addMove(-2, 2, c.isMoveToEat)
	addMove(2, -2, c.isMoveToEat)
	addMove(-2, -2, c.isMoveToEat)

	return moves
}

func (c Checker) GetAvailableMovesToEat(desk *Field, from Coordinate) []Coordinate {
	var moves []Coordinate
	addMove := func(dx, dy int, checker func(field *Field, from, to Coordinate) (bool, Coordinate)) {
		move := Coordinate{from.X + dx, from.Y + dy}
		isMove, _ := checker(desk, from, move)
		if isMove {
			moves = append(moves, move)
		}
	}

	addMove(2, 2, c.isMoveToEat)
	addMove(-2, 2, c.isMoveToEat)
	addMove(2, -2, c.isMoveToEat)
	addMove(-2, -2, c.isMoveToEat)

	return moves
}

func (c Checker) GetOwnerId() int {
	return c.OwnerId
}

func (c Checker) Move(desk *Field, from Coordinate, way []Coordinate) (bool, Coordinate) {
	var isCanBeMoved bool
	var foodPosition Coordinate

	vertical := desk.BordersRight.X
	if c.OwnerId == 1 {
		vertical = desk.BordersLeft.X
	}

	for i, to := range way {
		isCanBeMoved, foodPosition = c.isMoveToEat(desk, from, to)
		if i == 0 {
			isMoveWithoutEat, _ := c.isMoveWithoutEat(desk, from, to)
			if isMoveWithoutEat {
				desk.Move(from, to)
				if to.X == vertical {
					desk.RemoveWithOutBin(to)
					desk.Put(to, King{c.OwnerId})
				}
				return true, to
			}
			if !isCanBeMoved {
				return false, from
			}
		}
		if !isCanBeMoved {
			return true, from
		}
		desk.Remove(foodPosition)
		desk.Move(from, to)
		from = to
		if to.X == vertical {
			desk.RemoveWithOutBin(to)
			king := King{c.OwnerId}
			desk.Put(to, king)
			return king.moveOnlyToEat(desk, to, way[i+1:])
		}
	}

	return true, from

}

func (c Checker) IsMoveOne(desk *Field, from, to Coordinate) (bool, Coordinate) {
	isMoveWithFood, foodPosition := c.isMoveToEat(desk, from, to)
	if isMoveWithFood {
		return isMoveWithFood, foodPosition
	}
	isMoveWithOutFood, _ := c.isMoveWithoutEat(desk, from, to)
	return isMoveWithOutFood, foodPosition
}

func (c Checker) isMoveWithoutEat(desk *Field, from, to Coordinate) (bool, Coordinate) {
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

func (c Checker) isMoveToEat(desk *Field, from, to Coordinate) (bool, Coordinate) {
	foodPosition := Coordinate{
		(to.X + from.X) / 2,
		(to.Y + from.Y) / 2}

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
