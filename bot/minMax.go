package bot

import (
	"chekers/core"
	"math"
)

type minMax struct {
	level int
	body  tree
}

func (c minMax) analyzeField(field *core.Field, gamerId int) (core.Coordinate,
	[]core.Coordinate) {
	var from core.Coordinate
	var way []core.Coordinate
	c.body = tree{
		&node{
			nil,
			field.GetCopy(),
			core.Coordinate{},
			[]core.Coordinate{},
			gamerId,
			0}}

	c.body.Build(c.level * 2)
	from, way = c.body.GetBestMove()
	c.body = tree{}
	return from, way
}

type tree struct {
	root *node
}

func (c *tree) Build(level int) {
	c.root.createChilds(level, c.root.gamerId)
}

func (c *tree) GetBestMove() (core.Coordinate, []core.Coordinate) {
	for _, i := range c.root.childs {
		if i.score == c.root.score {
			return i.from, i.way
		}
	}
	return core.Coordinate{}, nil
}

func (c *tree) GetRandomMove(random Random) (core.Coordinate, []core.Coordinate) {
	if len(c.root.childs) != 0 {
		i := random.randn(len(c.root.childs))
		return c.root.childs[i].from, c.root.childs[i].way
	}
	return core.Coordinate{}, nil
}

type node struct {
	childs  []*node
	field   core.Field
	from    core.Coordinate
	way     []core.Coordinate
	gamerId int
	score   int
}

func (c *node) createChilds(n int, gamerId int) int {
	if n == 1 {
		return c.calculateScore(gamerId)
	}

	isEat := false
	for from, figure := range c.field.Figures {
		if figure.GetOwnerId() != c.gamerId {
			continue
		}
		moves := figure.GetAvailableMovesToEat(&c.field, from)
		if len(moves) != 0 {
			isEat = true
		}
		for _, to := range moves {
			childField := c.field.GetCopy()
			figure := childField.At(from)
			figure.Move(&childField, from, []core.Coordinate{to})
			c.childs = append(c.childs, &node{
				nil,
				childField,
				from,
				[]core.Coordinate{to},
				c.gamerId ^ 1,
				0})
		}
	}
	if isEat {
		for i := 0; i < len(c.childs); i++ {
			child := c.childs[i]
			figure := child.field.At(child.way[len(child.way)-1])
			moves := figure.GetAvailableMovesToEat(&child.field,
				child.way[len(child.way)-1])
			for _, to := range moves {
				childField := child.field.GetCopy()
				figure := childField.At(child.way[len(child.way)-1])
				figure.Move(&childField, child.way[len(child.way)-1],
					[]core.Coordinate{to})
				way := make([]core.Coordinate, len(child.way)+1)
				copy(way, child.way)
				way[len(child.way)] = to
				c.childs = append(c.childs, &node{
					nil,
					childField,
					child.from,
					way,
					c.gamerId ^ 1,
					0})
			}
		}
	} else {
		for from, figure := range c.field.Figures {
			if figure.GetOwnerId() != c.gamerId {
				continue
			}
			moves := figure.GetAvailableMoves(&c.field, from)
			for _, to := range moves {
				childField := c.field.GetCopy()
				figure := childField.At(from)
				figure.Move(&childField, from, []core.Coordinate{to})
				c.childs = append(c.childs, &node{
					nil,
					childField,
					from,
					[]core.Coordinate{to},
					c.gamerId ^ 1,
					0})
			}
		}
	}
	if c.childs == nil {
		return c.calculateScore(gamerId)
	}

	if c.gamerId != gamerId {
		min := math.MaxInt
		for _, child := range c.childs {
			min = int(math.Min(float64(child.createChilds(n-1, gamerId)), float64(min)))
			child.childs = nil
		}
		c.score = min
		return min
	} else {
		max := math.MinInt
		for _, child := range c.childs {
			max = int(math.Max(float64(child.createChilds(n-1, gamerId)), float64(max)))
			child.childs = nil
		}
		c.score = max
		return max
	}
}

func (c *node) calculateScore(gamerId int) int {
	ans := 0

	for _, figure := range c.field.Figures {
		if figure.GetOwnerId() == gamerId {
			ans += 1
		} else {
			ans -= 1
		}
	}
	c.score = ans
	return ans
}
