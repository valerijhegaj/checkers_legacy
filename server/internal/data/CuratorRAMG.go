package data

import (
	"errors"

	"checkers/core"
	"checkers/saveLoad"
	"checkers/server/internal/errorsStrings"
	"checkers/server/internal/game"
)

func NewCuratorRAMG() GameCurator {
	return &CuratorRAMG{}
}

type CuratorRAMG struct {
	game   map[int]game.Game
	gameID map[string]int

	maxGameID int
}

func (c *CuratorRAMG) NewGame(
	gameName, password string, settings game.Settings,
) error {
	_, ok := c.gameID[gameName]
	if ok {
		return errors.New(errorsStrings.GameAlreadyExist)
	}

	c.game[c.maxGameID] = game.NewGame(settings, password)
	c.gameID[gameName] = c.maxGameID

	c.maxGameID++
	return nil
}

func (c *CuratorRAMG) GetGame(
	token string, gameName string,
) (saveLoad.Save, error) {
	userID, err := GetGlobalStorage().GetUserID(token)
	if err != nil {
		return saveLoad.Save{}, err
	}
	gameID, ok := c.gameID[gameName]
	if !ok {
		return saveLoad.Save{}, errors.New(errorsStrings.NotFound)
	}
	game := c.game[gameID]
	return game.GetGame(userID)
}

func (c *CuratorRAMG) LoginGame(
	token string, gameName string,
) error {
	return nil
}
func (c *CuratorRAMG) MakeMove(
	token string, from core.Coordinate, path []core.Coordinate,
) error {
	return nil
}
func (c *CuratorRAMG) DeleteGame(
	token string, gameName string,
) error {
	return nil
}
