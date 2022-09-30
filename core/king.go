package core

type King struct {
	OwnerId int
}

func (c King) addMove(moves *[]Coordinate,
	desk *Field,
	d, from Coordinate,
	IsMove func(field *Field, from, to Coordinate) (bool, Coordinate)) {
	move := Coordinate{from.X + d.X, from.Y + d.Y}
	for ; desk.InBorders(move); move.X, move.Y = move.X+d.X, move.Y+d.Y {
		isMove, _ := IsMove(desk, from, move)
		if isMove {
			*moves = append(*moves, move)
		}
	}
}

func (c King) GetAvailableMoves(desk *Field, from Coordinate) []Coordinate {
	var moves []Coordinate
	c.addMove(&moves, desk, Coordinate{1, 1}, from, c.IsMoveOne)
	c.addMove(&moves, desk, Coordinate{1, -1}, from, c.IsMoveOne)
	c.addMove(&moves, desk, Coordinate{-1, 1}, from, c.IsMoveOne)
	c.addMove(&moves, desk, Coordinate{-1, -1}, from, c.IsMoveOne)
	return moves
}

func (c King) GetAvailableMovesToEat(desk *Field, from Coordinate) []Coordinate {
	var moves []Coordinate
	c.addMove(&moves, desk, Coordinate{1, 1}, from, c.isMoveOneToEat)
	c.addMove(&moves, desk, Coordinate{1, -1}, from, c.isMoveOneToEat)
	c.addMove(&moves, desk, Coordinate{-1, 1}, from, c.isMoveOneToEat)
	c.addMove(&moves, desk, Coordinate{-1, -1}, from, c.isMoveOneToEat)
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

func (c King) isMoveOneToEat(desk *Field, from, to Coordinate) (bool, Coordinate) {
	_, isWasFood, foodPosition := c.isMoveOne(desk, from, to)
	return isWasFood, foodPosition
}

func (c King) isMoveOne(desk *Field, from, to Coordinate) (bool, bool, Coordinate) {
	finishFoodPosition := desk.BordersLeft
	dx, dy := to.X-from.X, to.Y-from.Y
	if dx == 0 || (dx != dy && dx != -dy) || !desk.IsAvailable(to) {
		return false, false, desk.BordersLeft
	}

	var dx1, dy1 int
	if dx > 0 {
		dx1 = 1
	} else {
		dx1 = -1
	}
	if dy > 0 {
		dy1 = 1
	} else {
		dy1 = -1
	}

	wasAlreadyFood := false
	for i := 1; i < dx*dx1; i++ {
		foodPosition := Coordinate{from.X + i*dx1, from.Y + i*dy1}
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
