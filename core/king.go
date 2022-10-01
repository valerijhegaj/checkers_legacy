package core

type King struct {
	OwnerId int
}

func (c King) GetOwnerId() int {
	return c.OwnerId
}

func (c King) Move(
	desk *Field,
	from Coordinate,
	way []Coordinate,
) (
	bool,
	Coordinate,
) {
	var isCanBeMoved, isWasFood bool
	var foodPosition Coordinate

	isCanBeMoved, isWasFood, foodPosition = c.isMoveOne(
		desk, from, way[0],
	)
	if isCanBeMoved && !isWasFood {
		desk.Move(from, way[0])
		return true, way[0]
	}
	if !isCanBeMoved {
		return false, from
	}
	desk.Remove(foodPosition)
	desk.Move(from, way[0])

	return c.moveToEat(desk, way[0], way[1:])
}

// always returns true
func (c King) moveToEat(
	desk *Field,
	from Coordinate,
	way []Coordinate,
) (
	bool,
	Coordinate,
) {
	var isCanBeMoved, isWasFood bool
	var foodPosition Coordinate

	for _, to := range way {
		isCanBeMoved, isWasFood, foodPosition = c.isMoveOne(
			desk,
			from,
			to,
		)
		if !isCanBeMoved || !isWasFood {
			return true, from
		}
		desk.Remove(foodPosition)
		desk.Move(from, to)
		from = to
	}

	return true, from
}

func (c King) IsMoveOne(desk *Field, from, to Coordinate) (
	bool,
	Coordinate,
) {
	ans, _, foodPosition := c.isMoveOne(desk, from, to)
	return ans, foodPosition
}

func (c King) isMoveOneToEat(desk *Field, from, to Coordinate) (
	bool,
	Coordinate,
) {
	_, isWasFood, foodPosition := c.isMoveOne(desk, from, to)
	return isWasFood, foodPosition
}

func (c King) isMoveOne(desk *Field, from, to Coordinate) (
	bool,
	bool,
	Coordinate,
) {
	finishFoodPosition := desk.BordersLeft
	delta := Coordinate{to.X - from.X, to.Y - from.Y}
	if delta.X == 0 ||
		(delta.X != delta.Y && delta.X != -delta.Y) ||
		!desk.IsAvailable(to) {
		return false, false, desk.BordersLeft
	}

	direction := c.getDirection(delta)
	wasAlreadyFood := false
	foodPosition := Coordinate{
		from.X + direction.X, from.Y + direction.Y,
	}
	for ; foodPosition != to; foodPosition.X, foodPosition.Y =
		foodPosition.X+direction.X, foodPosition.Y+direction.Y {
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

func (c King) getDirection(delta Coordinate) Coordinate {
	var direction Coordinate
	if delta.X > 0 {
		direction.X = 1
	} else {
		direction.X = -1
	}
	if delta.Y > 0 {
		direction.Y = 1
	} else {
		direction.Y = -1
	}
	return direction
}

func (c King) addMove(
	availableMoves *[]Coordinate,
	desk *Field,
	direction, from Coordinate,
	IsMove func(field *Field, from, to Coordinate) (
		bool,
		Coordinate,
	),
) {
	move := Coordinate{from.X + direction.X, from.Y + direction.Y}
	for ; desk.InBorders(move); move.X, move.Y =
		move.X+direction.X, move.Y+direction.Y {
		isMove, _ := IsMove(desk, from, move)
		if isMove {
			*availableMoves = append(*availableMoves, move)
		}
	}
}

func (c King) GetAvailableMoves(
	desk *Field,
	from Coordinate,
) []Coordinate {
	var moves []Coordinate
	c.addMove(&moves, desk, Coordinate{1, 1}, from, c.IsMoveOne)
	c.addMove(&moves, desk, Coordinate{1, -1}, from, c.IsMoveOne)
	c.addMove(&moves, desk, Coordinate{-1, 1}, from, c.IsMoveOne)
	c.addMove(&moves, desk, Coordinate{-1, -1}, from, c.IsMoveOne)
	return moves
}

func (c King) GetAvailableMovesToEat(
	desk *Field,
	from Coordinate,
) []Coordinate {
	var moves []Coordinate
	c.addMove(&moves, desk, Coordinate{1, 1}, from, c.isMoveOneToEat)
	c.addMove(&moves, desk, Coordinate{1, -1}, from, c.isMoveOneToEat)
	c.addMove(&moves, desk, Coordinate{-1, 1}, from, c.isMoveOneToEat)
	c.addMove(&moves, desk, Coordinate{-1, -1}, from, c.isMoveOneToEat)
	return moves
}
