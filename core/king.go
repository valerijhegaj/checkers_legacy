package core

type King struct {
	OwnerId int
}

func (c King) GetOwnerId() int {
	return c.OwnerId
}

func (c King) Move(desk *Field, actualPosition Coordinate, newPosition []Coordinate) (bool, Coordinate) {
	deadNum := len(desk.Bin)

	for i, newPositionOne := range newPosition {
		if i > 0 && deadNum-len(desk.Bin) == 0 {
			return true, actualPosition
		} else {
			deadNum = len(desk.Bin)
		}
		isMoved, isWasFood := c.moveOne(desk, actualPosition, newPositionOne, true)
		if !isMoved {
			if i == 0 {
				return false, actualPosition
			} else {
				break
			}
		}
		if !isWasFood {
			//can optimize
			c.moveOne(desk, newPositionOne, actualPosition, true)
			break
		}
		actualPosition = newPositionOne
	}
	return true, actualPosition
}

func (c King) IsMoveOne(desk *Field, actualPosition, newPosition Coordinate) bool {
	ans, _ := c.moveOne(desk, actualPosition, newPosition, false)
	return ans
}

func (c King) moveOne(desk *Field, actualPosition, newPosition Coordinate, isMakeMove bool) (bool, bool) {
	dx, dy := newPosition.X-actualPosition.X, newPosition.Y-actualPosition.Y
	if dx == 0 || (dx != dy && dx != -dy) || !desk.IsFree(newPosition) {
		return false, false
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
	finishFoodPosition := Coordinate{}
	for i := 1; i < dx*dx_1; i++ {
		foodPosition := Coordinate{actualPosition.X + i*dx_1, actualPosition.Y + i*dy_1}
		if !desk.IsFree(foodPosition) {
			if wasAlreadyFood {
				return false, false
			}
			food := desk.At(foodPosition)
			if food.GetOwnerId() == c.GetOwnerId() {
				return false, false
			}
			finishFoodPosition = foodPosition
			wasAlreadyFood = true
		}
	}

	if isMakeMove {
		if wasAlreadyFood {
			desk.Remove(finishFoodPosition)
		}
		desk.Move(actualPosition, newPosition)
	}
	return true, wasAlreadyFood
}
