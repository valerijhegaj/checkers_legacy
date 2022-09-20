package core

type Figure interface {
	GetOwnerId() int
	// in each Figure implemented rules where you can move
	Move(desk *Field, actualPosition Coordinate, newPosition ...Coordinate) (bool, Coordinate)
}

type Checker struct {
	OwnerId int
}

func (c Checker) GetOwnerId() int {
	return c.OwnerId
}

func (c Checker) Move(desk *Field, actualPosition Coordinate, newPosition ...Coordinate) (bool, Coordinate) {
	for i, newPositionOne := range newPosition {
		if !c.moveOne(desk, actualPosition, newPositionOne) {
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

func (c Checker) moveOne(desk *Field, actualPosition Coordinate, newPosition Coordinate) bool {
	return c.moveWithoutEat(desk, actualPosition, newPosition) || c.moveToEat(desk, actualPosition, newPosition)
}

func (c Checker) moveWithoutEat(desk *Field, actualPosition Coordinate, newPosition Coordinate) bool {
	if newPosition.X-actualPosition.X == 1 &&
		(newPosition.Y-actualPosition.Y == 1 || newPosition.Y-actualPosition.Y == -1) {
		if desk.IsFree(newPosition) {
			desk.Move(actualPosition, newPosition)
			return true
		}
	}
	return false
}

func (c Checker) moveToEat(desk *Field, actualPosition Coordinate, newPosition Coordinate) bool {
	foodPosition := Coordinate{
		(newPosition.X + actualPosition.X) / 2,
		(newPosition.Y + actualPosition.Y) / 2}

	if (newPosition.X-actualPosition.X == 2 || newPosition.X-actualPosition.Y == -2) &&
		(newPosition.Y-actualPosition.Y == 2 || newPosition.Y-actualPosition.Y == -2) {
		if desk.IsFree(newPosition) && !desk.IsFree(foodPosition) {
			food := desk.At(foodPosition)
			if food.GetOwnerId() != c.GetOwnerId() {
				desk.Remove(foodPosition)
				desk.Move(actualPosition, newPosition)
				return true
			}
		}
	}

	return false
}

type King struct {
	OwnerId int
}

func (c King) GetOwnerId() int {
	return c.OwnerId
}

func (c King) Move(desk *Field, actualPosition Coordinate, newPosition ...Coordinate) (bool, Coordinate) {
	for i, newPositionOne := range newPosition {
		if !c.moveOne(desk, actualPosition, newPositionOne) {
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

func (c King) moveOne(desk *Field, actualPosition Coordinate, newPosition Coordinate) bool {
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

	if wasAlreadyFood {
		desk.Remove(finishFoodPosition)
	}
	desk.Move(actualPosition, newPosition)
	return true
}
