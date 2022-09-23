package core

type King struct {
	OwnerId int
}

func (c King) GetOwnerId() int {
	return c.OwnerId
}

func (c King) Move(desk *Field, from Coordinate, newPosition []Coordinate) (bool, Coordinate) {
	var isCanBeMoved, isWasFood bool
	var foodPosition Coordinate

	for i, to := range newPosition {
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
func (c King) moveOnlyToEat(desk *Field, from Coordinate, newPosition []Coordinate) (bool, Coordinate) {
	var isCanBeMoved, isWasFood bool
	var foodPosition Coordinate

	for _, to := range newPosition {
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

func (c King) IsMoveOne(desk *Field, actualPosition, newPosition Coordinate) bool {
	ans, _, _ := c.isMoveOne(desk, actualPosition, newPosition)
	return ans
}

func (c King) isMoveOne(desk *Field, actualPosition, newPosition Coordinate) (bool, bool, Coordinate) {
	finishFoodPosition := Coordinate{}
	dx, dy := newPosition.X-actualPosition.X, newPosition.Y-actualPosition.Y
	if dx == 0 || (dx != dy && dx != -dy) || !desk.IsAvailable(newPosition) {
		return false, false, finishFoodPosition
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
		foodPosition := Coordinate{actualPosition.X + i*dx_1, actualPosition.Y + i*dy_1}
		if !desk.IsAvailable(foodPosition) {
			if wasAlreadyFood {
				return false, false, foodPosition
			}
			food := desk.At(foodPosition)
			if food.GetOwnerId() == c.GetOwnerId() {
				return false, false, foodPosition
			}
			finishFoodPosition = foodPosition
			wasAlreadyFood = true
		}
	}

	return true, wasAlreadyFood, finishFoodPosition
}
