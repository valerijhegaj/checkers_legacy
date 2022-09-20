package core

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c *Coordinate) InitFromString(coordinate string) {
	c.X = int(coordinate[1] - '1')
	c.Y = int(coordinate[0] - 'a')
}
