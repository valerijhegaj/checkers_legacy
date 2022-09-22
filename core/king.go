package core

type King struct {
	OwnerId int
}

func (c King) GetOwnerId() int {
	return c.OwnerId
}

func (c King) Move(desk *Field, actualPosition Coordinate, newPosition []Coordinate) (bool, Coordinate) {
	for i, newPositionOne := range newPosition {
		if !c.moveOne(desk, actualPosition, newPositionOne, true) {
			if i == 0 {
				return false, actualPosition
			} else {
				break
			}
		}
		actualPosition = newPositionOne
	}
	return true, actualPosition
}

func (c King) IsMoveOne(desk *Field, actualPosition, newPosition Coordinate) bool {
	return c.moveOne(desk, actualPosition, newPosition, false)
}

func (c King) moveOne(desk *Field, actualPosition, newPosition Coordinate, isMakeMove bool) bool {
	dx, dy := newPosition.X-actualPosition.X, newPosition.Y-actualPosition.Y
	if dx == 0 && dx != dy && dx != -dy && !desk.IsFree(newPosition) {
		return false
	}

	wasAlreadyFood := false
	finishFoodPosition := Coordinate{}
	for i := 0; i < dx; i++ {
		foodPosition := Coordinate{actualPosition.X + i, actualPosition.Y + i*dy/dy}
		if !desk.IsFree(foodPosition) {
			if wasAlreadyFood {
				return false
			}
			food := desk.At(foodPosition)
			if food.GetOwnerId() == c.GetOwnerId() {
				return false
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
	return true
}
