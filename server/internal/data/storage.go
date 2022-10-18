package data

import (
	"checkers/core"
	"checkers/server/internal/game"
)

var GlobalStorage Storage

func InitGlobalStorage() error {
	GlobalStorage = struct {
		UserCurator
		GameCurator
	}{
		UserCurator: NewCuratorRAMU(),
		GameCurator: NewCuratorRAMG(),
	}
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
	NewGame(gameName, password string, settings game.Settings) error

	GetGame(token, gameName string) ([]byte, error)
	LoginGame(token, gameName, password string) error
	ChangeGame(token, gameName string, settings game.Settings) error
	DeleteGame(token, gameName string) error

	MakeMove(
		token, gameName string, from core.Coordinate,
		path []core.Coordinate,
	) error
}
