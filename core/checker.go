package core

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

func (c Checker) IsMoveOne(desk *Field, actualPosition, newPosition Coordinate) bool {
	return c.moveWithoutEat(desk, actualPosition, newPosition, false) ||
		c.moveToEat(desk, actualPosition, newPosition, false)
}

func (c Checker) moveOne(desk *Field, actualPosition, newPosition Coordinate) bool {
	return c.moveWithoutEat(desk, actualPosition, newPosition, true) ||
		c.moveToEat(desk, actualPosition, newPosition, true)
}

func (c Checker) moveWithoutEat(desk *Field, actualPosition, newPosition Coordinate, isMakeMove bool) bool {
	if newPosition.X-actualPosition.X == 1 &&
		(newPosition.Y-actualPosition.Y == 1 || newPosition.Y-actualPosition.Y == -1) {
		if desk.IsFree(newPosition) {
			if isMakeMove {
				desk.Move(actualPosition, newPosition)
			}
			return true
		}
	}
	return false
}

func (c Checker) moveToEat(desk *Field, actualPosition, newPosition Coordinate, isMakeMove bool) bool {
	foodPosition := Coordinate{
		(newPosition.X + actualPosition.X) / 2,
		(newPosition.Y + actualPosition.Y) / 2}

	if (newPosition.X-actualPosition.X == 2 || newPosition.X-actualPosition.Y == -2) &&
		(newPosition.Y-actualPosition.Y == 2 || newPosition.Y-actualPosition.Y == -2) {
		if desk.IsFree(newPosition) && !desk.IsFree(foodPosition) {
			food := desk.At(foodPosition)
			if food.GetOwnerId() != c.GetOwnerId() {
				if isMakeMove {
					desk.Remove(foodPosition)
					desk.Move(actualPosition, newPosition)
				}
				return true
			}
		}
	}

	return false
}
