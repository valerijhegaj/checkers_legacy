package data

import (
	"checkers/core"
	"checkers/saveLoad"
	"checkers/server/internal/game"
)

var GlobalStorage Storage

func InitGlobalStorage() error {
	GlobalStorage = struct {
		UserCurator
		GameCurator
	}{UserCurator: NewCuratorRAMU()}
	return nil
}

func GetGlobalStorage() Storage {
	return GlobalStorage
}

type Storage interface {
	UserCurator
	GameCurator
}

type UserCurator interface {
	NewUser(username, password string) error
	NewSession(username, password string) (string, error)

	GetUserID(token string) (int, error)
	GetUsername(userID int) (string, error)
}

type GameCurator interface {
	NewGame(token, gameName string, settings game.Settings) error

	GetGame(token, gameName string) (saveLoad.Save, error)
	LoginGame(token, gameName string) error
	MakeMove(
		token string, from core.Coordinate, path []core.Coordinate,
	) error
	DeleteGame(token, gameName string) error
}
