package bot

import (
	"chekers/core"
	"math/rand"
)

type RandomMoves struct {
	Random
}

func (c RandomMoves) analyzeField(field *core.Field, gamerId int) (core.Coordinate, []core.Coordinate) {
	for {
		from := c.getRandomPosition(field.BordersLeft, field.BordersRight)
		figure := field.At(from)
		if figure == nil {
			continue
		}
		moves := figure.GetAvailableMoves(field, from)
		if moves == nil {
			continue
		}
		to := moves[c.randn(len(moves))]
		return from, []core.Coordinate{to}
	}
}

func (c RandomMoves) getRandomPosition(left, right core.Coordinate) core.Coordinate {
	return core.Coordinate{c.randlr(left.X, right.X+1), c.randlr(left.Y, right.Y+1)}
}

type Random interface {
	randlr(l, t int) int
	randn(n int) int
}

type PsevdoRandom struct{}

func (c PsevdoRandom) randlr(l, r int) int {
	return rand.Intn(r-l) + l
}

func (c PsevdoRandom) randn(n int) int {
	return rand.Intn(n)
}
