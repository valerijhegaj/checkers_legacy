package apiParser

import (
	"fmt"
	"net/http"

	"checkers/server/api"
	"checkers/server/pkg/file"
)

type User struct {
	Username, Password string
	cookies            []*http.Cookie

	client http.Client
	PORT   int
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
