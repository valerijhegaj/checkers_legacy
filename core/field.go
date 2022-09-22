package core

// non sync
type Field struct {
	Figures      map[Coordinate]Figure
	Bin          []Figure
	BordersRight Coordinate
	BordersLeft  Coordinate
}

func (c *Field) Init() {
	c.Figures = make(map[Coordinate]Figure)
}

func (c *Field) InBorders(coordinate Coordinate) bool {
	return coordinate.X <= c.BordersRight.X && coordinate.Y <= c.BordersRight.Y &&
		coordinate.X >= c.BordersLeft.X && coordinate.Y >= c.BordersLeft.Y
}

func (c *Field) IsFree(coordinate Coordinate) bool {
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

func (c *Field) Move(prev Coordinate, next Coordinate) {
	if !c.IsFree(prev) {
		c.Figures[next] = c.Figures[prev]
		delete(c.Figures, prev)
	}
}

func (c *Field) Remove(ptr Coordinate) {
	c.Bin = append(c.Bin, c.Figures[ptr])
	delete(c.Figures, ptr)
}

func (c *Field) Put(ptr Coordinate, figure Figure) {
	c.Figures[ptr] = figure
}

func (c *Field) GetBin() []Figure {
	return c.Bin
}

func (c *Field) GetCopy() Field {
	var copy_ Field
	copy_.Init()

	for key, value := range c.Figures {
		copy_.Figures[key] = value
	}

	copy_.Bin = make([]Figure, len(c.Bin))
	copy(copy_.Bin, c.Bin)

	copy_.BordersRight = c.BordersRight
	copy_.BordersLeft = c.BordersLeft

	return copy_
}
