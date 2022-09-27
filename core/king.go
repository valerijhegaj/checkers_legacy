package core

type King struct {
	OwnerId int
}

func (c King) GetAvailableMoves(desk *Field, from Coordinate) []Coordinate {
	var moves []Coordinate
	addMoves := func(dx, dy int, checker func(field *Field, from, to Coordinate) (bool, Coordinate)) {
		move := Coordinate{from.X + dx, from.Y + dy}
		for ; desk.InBorders(move); move.X, move.Y = move.X+1, move.Y+1 {
			isMove, _ := checker(desk, from, move)
			if isMove {
				moves = append(moves, move)
			}
		}
	}

	addMoves(1, 1, c.IsMoveOne)
	addMoves(1, -1, c.IsMoveOne)
	addMoves(-1, 1, c.IsMoveOne)
	addMoves(-1, -1, c.IsMoveOne)

	return moves
}

func (c King) GetOwnerId() int {
	return c.OwnerId
}

func (c King) Move(desk *Field, from Coordinate, way []Coordinate) (bool, Coordinate) {
	var isCanBeMoved, isWasFood bool
	var foodPosition Coordinate

	for i, to := range way {
		isCanBeMoved, isWasFood, foodPosition = c.isMoveOne(desk, from, to)
		if i == 0 {
			if isCanBeMoved && !isWasFood {
				desk.Move(from, to)
				return true, to
			}
			if !isCanBeMoved {
				return false, from
			}
		}
		if !isCanBeMoved || !isWasFood {
			return true, from
		}
		desk.Remove(foodPosition)
		desk.Move(from, to)
		from = to
	}

	return true, from
}

// always returns true, method for checker (test in checker)
func (c King) moveOnlyToEat(desk *Field, from Coordinate, way []Coordinate) (bool, Coordinate) {
	var isCanBeMoved, isWasFood bool
	var foodPosition Coordinate

	for _, to := range way {
		isCanBeMoved, isWasFood, foodPosition = c.isMoveOne(desk, from, to)
		if !isCanBeMoved || !isWasFood {
			return true, from
		}
		desk.Remove(foodPosition)
		desk.Move(from, to)
		from = to
	}

	return true, from
}

func (c King) IsMoveOne(desk *Field, from, to Coordinate) (bool, Coordinate) {
	ans, _, foodPosition := c.isMoveOne(desk, from, to)
	return ans, foodPosition
}

func (c King) isMoveOne(desk *Field, from, to Coordinate) (bool, bool, Coordinate) {
	finishFoodPosition := desk.BordersLeft
	dx, dy := to.X-from.X, to.Y-from.Y
	if dx == 0 || (dx != dy && dx != -dy) || !desk.IsAvailable(to) {
		return false, false, desk.BordersLeft
	}

	var dx_1, dy_1 int
	if dx > 0 {
		dx_1 = 1
	} else {
		dx_1 = -1
	}
	if dy > 0 {
		dy_1 = 1
	} else {
		dy_1 = -1
	}

	wasAlreadyFood := false
	for i := 1; i < dx*dx_1; i++ {
		foodPosition := Coordinate{from.X + i*dx_1, from.Y + i*dy_1}
		if !desk.IsAvailable(foodPosition) {
			if wasAlreadyFood {
				return false, false, desk.BordersLeft
			}
			food := desk.At(foodPosition)
			if food.GetOwnerId() == c.GetOwnerId() {
				return false, false, desk.BordersLeft
			}
			finishFoodPosition = foodPosition
			wasAlreadyFood = true
		}
	}

	return true, wasAlreadyFood, finishFoodPosition
}
