package bot

import (
	"chekers/core"
	"math/rand"
)

func CreateRandomMoves() RandomMoves {
	return RandomMoves{psevdoRandom{}, tree{}}
}

type RandomMoves struct {
	Random
	body tree
}

func (c RandomMoves) analyzeField(field *core.Field, gamerId int) (core.Coordinate, []core.Coordinate) {
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

	c.body.Build(2)
	from, way = c.body.GetRandomMove(c.Random)
	c.body = tree{}
	return from, way
}

func (c RandomMoves) getRandomPosition(left, right core.Coordinate) core.Coordinate {
	return core.Coordinate{c.randlr(left.X, right.X+1), c.randlr(left.Y, right.Y+1)}
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
