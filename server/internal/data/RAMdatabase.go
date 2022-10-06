package data

import "errors"

func NewRAMstorage() Storage {
	return &RAMstorage{&RAMdatabase}
}

var RAMdatabase dataBase

type dataBase struct {
	users  map[string]string
	games  map[string]Game
	tokens map[string]string
}

type Game struct{}

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

func (c *RAMstorage) DeleteUser(name string) {
	delete(c.users, name)
}

func (c *RAMstorage) ChangePassword(token, password string) error {
	_, ok := c.tokens[token]
	if !ok {
		return errors.New(ErrorBadToken)
	}

	c.users[c.tokens[token]] = password
	return nil
}
