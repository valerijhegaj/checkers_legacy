package data

import (
	"errors"

	"checkers/server/internal/errorsStrings"
)

func NewCuratorRAMU() UserCurator {
	return &CuratorRAMU{
		users:     make(map[int][2]string),
		userID:    make(map[string]int),
		token:     make(map[string]int),
		maxUserID: 1,
	}
}

type CuratorRAMU struct {
	owner *Storage

	users  map[int][2]string
	userID map[string]int
	token  map[string]int

	maxUserID int
}

func (c *CuratorRAMU) Init(owner *Storage) {
	c.owner = owner
}
func (c *CuratorRAMU) NewUser(
	username string, password string,
) error {
	if _, userAlreadyExist := c.userID[username]; userAlreadyExist {
		return errors.New(errorsStrings.UserAlreadyExist)
	}

	c.users[c.maxUserID] = [2]string{username, password}
	c.userID[username] = c.maxUserID
	c.maxUserID++
	return nil
}

func (c *CuratorRAMU) NewSession(
	username string, password string,
) (string, error) {
	userID, isExistUser := c.userID[username]
	if !isExistUser {
		return "wrong name", errors.New(errorsStrings.PermissionDenied)
	}
	if c.users[userID][1] != password {
		return "wrong password", errors.New(errorsStrings.PermissionDenied)
	}
	//generate token
	token := username

	c.token[token] = userID
	return token, nil
}

func (c *CuratorRAMU) GetUserID(token string) (int, error) {
	userID, ok := c.token[token]
	if !ok {
		return 0, errors.New(errorsStrings.NotFound)
	}
	return userID, nil
}

func (c *CuratorRAMU) GetUsername(userID int) (string, error) {
	user, ok := c.users[userID]
	if !ok {
		return "", errors.New(errorsStrings.NotFound)
	}
	return user[0], nil
}
