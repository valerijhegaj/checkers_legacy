package core

func NewField() Field {
	return Field{Figures: make(map[Coordinate]Figure)}
}

func NewStandart8x8Field() Field {
	field := NewField()

	field.Put(Coordinate{0, 0}, Checker{0})
	field.Put(Coordinate{0, 2}, Checker{0})
	field.Put(Coordinate{0, 4}, Checker{0})
	field.Put(Coordinate{0, 6}, Checker{0})
	field.Put(Coordinate{1, 1}, Checker{0})
	field.Put(Coordinate{1, 3}, Checker{0})
	field.Put(Coordinate{1, 5}, Checker{0})
	field.Put(Coordinate{1, 5}, Checker{0})
	field.Put(Coordinate{1, 7}, Checker{0})
	field.Put(Coordinate{2, 0}, Checker{0})
	field.Put(Coordinate{2, 2}, Checker{0})
	field.Put(Coordinate{2, 4}, Checker{0})
	field.Put(Coordinate{2, 6}, Checker{0})

	field.Put(Coordinate{5, 1}, Checker{1})
	field.Put(Coordinate{5, 3}, Checker{1})
	field.Put(Coordinate{5, 5}, Checker{1})
	field.Put(Coordinate{5, 7}, Checker{1})
	field.Put(Coordinate{6, 0}, Checker{1})
	field.Put(Coordinate{6, 2}, Checker{1})
	field.Put(Coordinate{6, 4}, Checker{1})
	field.Put(Coordinate{6, 6}, Checker{1})
	field.Put(Coordinate{7, 1}, Checker{1})
	field.Put(Coordinate{7, 3}, Checker{1})
	field.Put(Coordinate{7, 5}, Checker{1})
	field.Put(Coordinate{7, 7}, Checker{1})

	field.BordersRight = Coordinate{7, 7}
	return field
}

type Field struct {
	Figures      map[Coordinate]Figure
	Bin          []Figure
	BordersRight Coordinate
	BordersLeft  Coordinate
}

func (c *Field) InBorders(coordinate Coordinate) bool {
	return coordinate.X <= c.BordersRight.X &&
		coordinate.Y <= c.BordersRight.Y &&
		coordinate.X >= c.BordersLeft.X &&
		coordinate.Y >= c.BordersLeft.Y
}

func (c *Field) IsAvailable(coordinate Coordinate) bool {
	_, ok := c.Figures[coordinate]
	return !ok && c.InBorders(coordinate)
}

func (c *Field) At(coordinate Coordinate) Figure {
	ans, ok := c.Figures[coordinate]
	if ok {
		return ans
	}
	return nil
}

func (c *Field) Move(from Coordinate, to Coordinate) {
	if !c.IsAvailable(from) {
		c.Figures[to] = c.Figures[from]
		delete(c.Figures, from)
	}
}

func (c *Field) Remove(ptr Coordinate) {
	c.Bin = append(c.Bin, c.Figures[ptr])
	c.RemoveWithOutBin(ptr)
}

func (c *Field) RemoveWithOutBin(ptr Coordinate) {
	delete(c.Figures, ptr)
}

func (c *Field) Put(ptr Coordinate, figure Figure) {
	c.Figures[ptr] = figure
}

func (c *Field) GetCopy() Field {
	copy_ := NewField()

	for key, value := range c.Figures {
		copy_.Figures[key] = value
	}

	copy_.Bin = make([]Figure, len(c.Bin))
	copy(copy_.Bin, c.Bin)

	copy_.BordersRight = c.BordersRight
	copy_.BordersLeft = c.BordersLeft

	return copy_
}

func (c *Field) CreateStandart88() {

}
