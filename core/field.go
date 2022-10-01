package core

func NewField() Field {
	return Field{Figures: make(map[Coordinate]Figure)}
}

func NewStandartField(n int) Field {
	field := NewField()
	for x := 0; x < n/2-1; x++ {
		for y := x % 2; y < n; y += 2 {
			field.Put(Coordinate{x, y}, Checker{0})
		}
	}

	for x := n/2 + 1; x < n; x++ {
		for y := x % 2; y < n; y += 2 {
			field.Put(Coordinate{x, y}, Checker{1})
		}
	}

	field.BordersRight = Coordinate{n - 1, n - 1}
	return field
}

func NewStandart8x8Field() Field {
	return NewStandartField(8)
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
