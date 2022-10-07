package logInGame

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"checkers/server/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(
			"bad method for login game, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, gameID, password, err := Parse(r.Body)
	if err != nil {
		log.Println("Tried to login game, but " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := data.GetStorage()
	if err != nil {
		log.Println("Tried to login game, but " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = d.LogInGame(token, gameID, password)
	if err != nil {
		log.Println("Tried to login game, but " + err.Error())
		if err.Error() == data.ErrorNotFoundGame {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == data.ErrorNotHaveAccess {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	log.Println("Log in gameID: " + gameID + ", token: " + token)
}

type helperParse struct {
	Token    string `json:"token"`
	GameID   string `json:"game_id"`
	Password string `json:"password"`
}

func Parse(i io.ReadCloser) (
	string,
	string,
	string,
	error,
) {
	data := make([]byte, 1024)
	n, err := i.Read(data)
	if err != nil && err != io.EOF {
		return "", "", "", err
	}

	var helper helperParse
	err = json.Unmarshal(data[:n], &helper)
	if err != nil {
		return "", "", "", err
	}
	return helper.Token, helper.GameID, helper.Password, nil
}
