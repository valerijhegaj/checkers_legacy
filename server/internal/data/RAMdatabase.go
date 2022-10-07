package data

import (
	"errors"

	"checkers/core"
	"checkers/saveLoad"
)

func NewRAMstorage() Storage {
	return &RAMstorage{&RAMdatabase}
}

var RAMdatabase dataBase

type dataBase struct {
	users  map[string]string
	games  map[string]Game
	tokens map[string]string
}

type RAMstorage struct {
	*dataBase
}

func (c *RAMstorage) Init() error {
	c.users = make(map[string]string)
	c.games = make(map[string]Game)
	c.tokens = make(map[string]string)
	return nil
}

func (c *RAMstorage) NewUser(name, password string) error {
	_, ok := c.users[name]
	if ok {
		return errors.New(ErrorAlreadyExist)
	}
	c.users[name] = password
	return nil
}

func (c *RAMstorage) NewToken(token, name string) error {
	c.tokens[token] = name
	return nil
}

func (c *RAMstorage) CheckAccess(name, password string) error {
	pass, ok := c.users[name]
	if !ok {
		return errors.New(ErrorWrongUserName)
	}
	if pass != password {
		return errors.New(ErrorWrongPassword)
	}
	return nil
}

func (c *RAMstorage) DeleteUser(token string) {
	delete(c.users, c.tokens[token])
	delete(c.tokens, token)
}

func (c *RAMstorage) ChangePassword(token, password string) error {
	_, ok := c.tokens[token]
	if !ok {
		return errors.New(ErrorBadToken)
	}

	c.users[c.tokens[token]] = password
	return nil
}

func (c *RAMstorage) CheckToken(token string) error {
	_, ok := c.tokens[token]
	if !ok {
		return errors.New(ErrorBadToken)
	}
	return nil
}

func (c *RAMstorage) NewGame(
	save saveLoad.Save,
	password string,
	gamerID int,
	token string,
) (
	string,
	error,
) {
	game := NewGame(save, password, gamerID, c.tokens[token])
	gameID := c.tokens[token]
	c.games[gameID] = game
	return gameID, nil
}

func (c *RAMstorage) LogOutGame(token, gameID string) error {
	game, ok := c.games[gameID]
	if !ok {
		return errors.New(ErrorNotFoundGame)
	} else if game.user[0] == c.tokens[token] {
		game.user[0] = ""
		if game.user[1] == "" {
			delete(c.games, gameID)
		}
		return nil
	} else if game.user[1] == c.tokens[token] {
		game.user[1] = ""
		if game.user[0] == "" {
			delete(c.games, gameID)
		}
		return nil
	}
	return errors.New(ErrorNotHaveAccess)
}

func (c *RAMstorage) LogInGame(
	token,
	gameID,
	password string,
) error {
	game, ok := c.games[gameID]
	if !ok {
		return errors.New(ErrorNotFoundGame)
	}
	if game.password != password {
		return errors.New(ErrorNotHaveAccess)
	}
	if game.user[0] == "" && game.Gamer0 != saveLoad.Bot {
		game.user[0] = c.tokens[token]
		c.games[gameID] = game
		return nil
	}
	if game.user[1] == "" && game.Gamer1 != saveLoad.Bot {
		game.user[1] = c.tokens[token]
		c.games[gameID] = game
		return nil
	}
	return errors.New(ErrorNotHaveAccess)
}

func (c *RAMstorage) Move(
	token, gameID string, from core.Coordinate,
	way []core.Coordinate,
) error {
	game, ok := c.games[gameID]
	if !ok {
		return errors.New(ErrorNotFoundGame)
	}
	name := c.tokens[token]
	if game.user[0] == name {
		success := game.Move(from, way, 0)
		if success {
			return nil
		}
	}
	if game.user[1] == name {
		success := game.Move(from, way, 1)
		if success {
			return nil
		}
	}

	return errors.New(ErrorNotHaveAccess)
}

func (c *RAMstorage) GetGame(
	token,
	gameID,
	password string,
) (
	saveLoad.Save,
	error,
) {
	name := c.tokens[token]
	game, ok := c.games[gameID]
	if !ok {
		return saveLoad.Save{}, errors.New(ErrorNotFoundGame)
	}
	if name == game.user[0] || name == game.user[1] {
		return game.GetSave(), nil
	}
	if game.password == password {
		return game.GetSave(), nil
	}
	return saveLoad.Save{}, errors.New(ErrorNotHaveAccess)
}
