package core

type Checker struct {
	OwnerId int
}

func (c Checker) GetOwnerId() int {
	return c.OwnerId
}

func (c Checker) Move(desk *Field, from Coordinate, newPosition []Coordinate) (bool, Coordinate) {
	var isCanBeMoved bool
	var foodPosition Coordinate

	vertical := desk.BordersRight.X
	if c.OwnerId == 1 {
		vertical = desk.BordersLeft.X
	}

	for i, to := range newPosition {
		isCanBeMoved, foodPosition = c.isMoveToEat(desk, from, to)
		if i == 0 {
			if c.isMoveWithoutEat(desk, from, to) {
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
			return king.moveOnlyToEat(desk, to, newPosition[i+1:])
		}
	}

	return true, from

}

func (c Checker) IsMoveOne(desk *Field, from, to Coordinate) bool {
	var isMoveWithFood bool
	isMoveWithFood, _ = c.isMoveToEat(desk, from, to)
	return isMoveWithFood || c.isMoveWithoutEat(desk, from, to)
}

func (c Checker) isMoveWithoutEat(desk *Field, from, to Coordinate) bool {
	vertical := 1
	if c.GetOwnerId() == 1 {
		vertical = -1
	}
	if to.X-from.X == vertical &&
		(to.Y-from.Y == 1 || to.Y-from.Y == -1) {
		if desk.IsAvailable(to) {
			return true
		}
	}
	return false
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

	return false, foodPosition
}
