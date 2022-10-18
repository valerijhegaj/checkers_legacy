package apiParser

import (
	"fmt"
	"io"
	"net/http"

	"checkers/core"
	"checkers/server/api"
	"checkers/server/pkg/defines"
	"checkers/server/pkg/file"
)

type User struct {
	Username, Password string
	cookies            []*http.Cookie

	client http.Client
	PORT   int
}

func (c *User) addCookies(req *http.Request) {
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}
}

func (c *User) IsEmptyCookies() bool {
	return len(c.cookies) == 0
}

func (c *User) Register() (int, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/user", c.PORT),
		file.NewReadCloserFromBytes(
			api.UnParse(
				api.Helper{
					UserName: c.Username, Password: c.Password,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}

	res, err := c.client.Do(req)
	return res.StatusCode, err
}

func (c *User) LogIn(maxAge int) (int, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/session", c.PORT),
		file.NewReadCloserFromBytes(
			api.UnParse(
				api.Helper{
					UserName: c.Username, Password: c.Password, MaxAge: maxAge,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}

	res, err := c.client.Do(req)

	if res.StatusCode == http.StatusCreated {
		c.cookies = res.Cookies()
	}

	return res.StatusCode, err
}

func (c *User) CreateGame(
	gameName, password string, settings defines.Settings,
) (int, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/game/create", c.PORT),
		file.NewReadCloserFromBytes(
			api.UnParse(
				api.Helper{
					GameName: gameName,
					Password: password,
					Settings: settings,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}

	c.addCookies(req)

	res, err := c.client.Do(req)
	return res.StatusCode, err
}

func (c *User) Move(
	gameName string, from core.Coordinate, path []core.Coordinate,
) (int, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/game/move", c.PORT),
		file.NewReadCloserFromBytes(
			api.UnParse(
				api.Helper{
					GameName: gameName,
					From:     from,
					Way:      path,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}

	c.addCookies(req)

	res, err := c.client.Do(req)
	return res.StatusCode, err
}

func (c *User) GetGame(gameName string) (int, []byte, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://localhost:%d/api/game", c.PORT),
		file.NewReadCloserFromBytes(
			api.UnParse(
				api.Helper{
					GameName: gameName,
				},
			),
		),
	)
	if err != nil {
		return -1, nil, err
	}

	c.addCookies(req)

	res, err := c.client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	if res.StatusCode == http.StatusOK {
		data, err := io.ReadAll(res.Body)
		return res.StatusCode, data, err
	}
	return res.StatusCode, nil, err
}

func (c *User) LogInGame(gameName, password string) (int, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/game", c.PORT),
		file.NewReadCloserFromBytes(
			api.UnParse(
				api.Helper{
					GameName: gameName,
					Password: password,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}

	c.addCookies(req)

	res, err := c.client.Do(req)
	return res.StatusCode, err
}
