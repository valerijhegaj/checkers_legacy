package data

import (
	"checkers/core"
	"checkers/saveLoad"
)

func InitStorage() error {
	storage, err := GetStorage()
	if err != nil {
		return err
	}
	storage.Init()
	return nil
}

func GetStorage() (
	Storage,
	error,
) {
	return NewRAMstorage(), nil
}

type Storage interface {
	NewUser(name, password string) error
	NewToken(token, name string) error
	ChangePassword(token, password string) error
	CheckAccess(name, password string) error
	CheckToken(token string) error
	DeleteUser(token string)

	NewGame(
		save saveLoad.Save, password string, gamerID int,
		token string,
	) (
		string,
		error,
	)
	LogOutGame(token, gameID string) error
	LogInGame(token, gameID, password string) error
	GetGame(token, gameID, password string) (
		saveLoad.Save,
		error,
	)
	Move(
		token, gameID string, from core.Coordinate,
		way []core.Coordinate,
	) error

	Init() error
}

const (
	ErrorBadToken      = "bad token"
	ErrorAlreadyExist  = "already exist"
	ErrorWrongUserName = "wrong username"
	ErrorWrongPassword = "wrong password"
	ErrorNotFoundGame  = "not found game"
	ErrorNotHaveAccess = "don't have access"
)
