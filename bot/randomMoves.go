package bot

import (
	"math/rand"

	"checkers/core"
)

func NewRandomMoves() RandomMoves {
	return RandomMoves{psevdoRandom{}, MinMaxTree{}}
}

type RandomMoves struct {
	Random
	body MinMaxTree
}

func (c RandomMoves) GetMove(
	field *core.Field,
	gamerId int,
) (
	core.Coordinate,
	[]core.Coordinate,
) {
	var from core.Coordinate
	var way []core.Coordinate
	c.body = MinMaxTree{2, nil, &simpleAmmount{}}
	from, way = c.body.GetRandomMove(c.Random)
	c.body = MinMaxTree{}
	return from, way
}

func (c RandomMoves) getRandomPosition(left, right core.Coordinate) core.Coordinate {
	return core.Coordinate{
		c.randlr(left.X, right.X+1),
		c.randlr(left.Y, right.Y+1),
	}
}

type Random interface {
	randlr(l, t int) int
	randn(n int) int
}

type psevdoRandom struct{}

func (c psevdoRandom) randlr(l, r int) int {
	return rand.Intn(r-l) + l
}

func (c psevdoRandom) randn(n int) int {
	return rand.Intn(n)
}
