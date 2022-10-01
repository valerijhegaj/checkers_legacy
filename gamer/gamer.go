package gamer

import (
	"chekers/core"
	"chekers/saveLoad"
)

type Gamer struct {
	GamerId int
	Core    *core.GameCore
}

func (c Gamer) GetField() core.Field {
	return c.Core.GetField()
}

func (c Gamer) IsTurn() bool {
	return c.Core.IsTurn(c.GamerId)
}

func (c Gamer) Move(
	from core.Coordinate,
	way []core.Coordinate,
) bool {
	ans := c.Core.Move(from, way, c.GamerId)
	//if ans {
	//	fmt.Println(from, way)
	//}
	return ans
}

func (c Gamer) InitSave(save saveLoad.Save) {
	c.Core.InitField(save.Field)
	c.Core.InitTurnGamerId(save.TurnGamerId)
}

func (c Gamer) GetWinner() (
	bool,
	Gamer,
) {
	field := c.GetField()
	var isCanMakeTurn [2]bool
	var numberFigures [2]int
	for from, figure := range field.Figures {
		numberFigures[figure.GetOwnerId()]++
		moves := figure.GetAvailableMoves(&field, from)
		if moves != nil {
			isCanMakeTurn[figure.GetOwnerId()] = true
		}
	}
	if numberFigures[0] == 0 || (!isCanMakeTurn[0] && c.Core.IsTurn(0)) {
		return true, Gamer{1, c.Core}
	}
	if numberFigures[1] == 0 || (!isCanMakeTurn[1] && c.Core.IsTurn(1)) {
		return true, Gamer{0, c.Core}
	}
	//if !isCanMakeTurn[0] && c.Core.IsTurn(0) || !isCanMakeTurn[1] && c.Core.IsTurn(1) {
	//	return true, Gamer{0, nil}
	//}
	return false, Gamer{0, nil}
}
