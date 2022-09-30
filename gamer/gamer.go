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

func (c Gamer) Move(from core.Coordinate, way []core.Coordinate) bool {
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

func (c Gamer) GetWinner() (bool, Gamer) {
	field := c.GetField()
	isCanMakeTurn0 := false
	isCanMakeTurn1 := false
	for from, figure := range field.Figures {
		moves := figure.GetAvailableMoves(&field, from)
		if moves != nil {
			if figure.GetOwnerId() == 0 {
				isCanMakeTurn0 = true
			} else {
				isCanMakeTurn1 = true
			}
		}
	}
	if (isCanMakeTurn1 || isCanMakeTurn0) == false {
		return true, Gamer{0, nil}
	} else if !isCanMakeTurn0 && c.Core.IsTurn(0) {
		return true, Gamer{1, c.Core}
	} else if !isCanMakeTurn1 && c.Core.IsTurn(1) {
		return true, Gamer{0, c.Core}
	}
	return false, Gamer{0, nil}
}
