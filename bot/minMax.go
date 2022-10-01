package bot

import (
	"math"
	"reflect"

	"chekers/core"
)

func NewMinMax(level int) Mind {
	return &MinMaxTree{
		level,
		nil,
		NewSimpleAmmount(),
	}
}

func NewMinMaxV2(level int, kingCost, checkerCost float64) Mind {
	return &MinMaxTree{
		level,
		nil,
		NewAmmountWithCoef(kingCost, checkerCost),
	}
}

type MinMaxTree struct {
	level int
	root  *node
	Evristika
}

func (c *MinMaxTree) GetMove(
	field *core.Field,
	gamerId int,
) (
	core.Coordinate,
	[]core.Coordinate,
) {
	c.root = &node{
		nil,
		*field,
		core.Coordinate{},
		[]core.Coordinate{},
		gamerId,
		0,
	}
	c.root.createChilds(c.level*2, c.root.gamerId, c.Evristika)
	defer func() { c.root = nil }()
	for _, i := range c.root.childs {
		if i.score == c.root.score {
			return i.from, i.way
		}
	}
	return core.Coordinate{}, nil
}

func (c *MinMaxTree) GetRandomMove(random Random) (
	core.Coordinate,
	[]core.Coordinate,
) {
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
	score   float64
}

func (c *node) createChilds(
	n int,
	gamerId int,
	evristika Evristika,
) float64 {
	if n == 1 {
		c.score = evristika.CalculateScore(gamerId, c.field)
		return c.score
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
			c.childs = append(
				c.childs, &node{
					nil,
					childField,
					from,
					[]core.Coordinate{to},
					c.gamerId ^ 1,
					0,
				},
			)
		}
	}
	if isEat {
		for i := 0; i < len(c.childs); i++ {
			child := c.childs[i]
			figure := child.field.At(child.way[len(child.way)-1])
			moves := figure.GetAvailableMovesToEat(
				&child.field,
				child.way[len(child.way)-1],
			)
			for _, to := range moves {
				childField := child.field.GetCopy()
				figure := childField.At(child.way[len(child.way)-1])
				figure.Move(
					&childField, child.way[len(child.way)-1],
					[]core.Coordinate{to},
				)
				way := make([]core.Coordinate, len(child.way)+1)
				copy(way, child.way)
				way[len(child.way)] = to
				c.childs = append(
					c.childs, &node{
						nil,
						childField,
						child.from,
						way,
						c.gamerId ^ 1,
						0,
					},
				)
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
				c.childs = append(
					c.childs, &node{
						nil,
						childField,
						from,
						[]core.Coordinate{to},
						c.gamerId ^ 1,
						0,
					},
				)
			}
		}
	}
	if c.childs == nil {
		c.score = evristika.CalculateScore(gamerId, c.field)
		return c.score
	}

	if c.gamerId != gamerId {
		min := math.MaxFloat64
		for _, child := range c.childs {
			min = math.Min(child.createChilds(n-1, gamerId, evristika), min)
			child.childs = nil
		}
		c.score = min
		return min
	} else {
		max := -math.MaxFloat64
		for _, child := range c.childs {
			max = math.Max(child.createChilds(n-1, gamerId, evristika), max)
			child.childs = nil
		}
		c.score = max
		return max
	}
}

func (c *node) calculateScore(gamerId int) float64 {
	ans := float64(0)

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

type Evristika interface {
	CalculateScore(gamerId int, field core.Field) float64
}

func NewSimpleAmmount() Evristika {
	return &simpleAmmount{}
}

type simpleAmmount struct {
}

func (c *simpleAmmount) CalculateScore(
	gamerId int,
	field core.Field,
) float64 {
	ans := float64(0)

	for _, figure := range field.Figures {
		if figure.GetOwnerId() == gamerId {
			ans += 1
		} else {
			ans -= 1
		}
	}
	return ans
}

func NewAmmountWithCoef(kingCost, checkerCost float64) Evristika {
	return &AmmountWithCoef{kingCost, checkerCost}
}

type AmmountWithCoef struct {
	KingCost    float64
	CheckerCost float64
}

func (c *AmmountWithCoef) CalculateScore(
	gamerId int,
	field core.Field,
) float64 {
	ans := float64(0)

	for _, figure := range field.Figures {
		var coef float64
		if reflect.TypeOf(figure) == reflect.TypeOf(core.Checker{}) {
			coef = c.CheckerCost
		} else {
			coef = c.KingCost
		}
		if figure.GetOwnerId() == gamerId {
			ans += coef
		} else {
			ans -= coef
		}
	}
	return ans
}
