package createGame

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"checkers/core"
	"checkers/saveLoad"
	"checkers/server/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(
			"Bad method for create game, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, password, partispants, gamerID, err := Parse(r.Body)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := data.GetStorage()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = d.CheckToken(token)
	if err != nil {
		log.Println(
			"Tried to create game token:" + token + ", " +
				"but " + err.Error(),
		)
		if err.Error() == data.ErrorBadToken {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	var save saveLoad.Save
	save.Field = core.NewStandard8x8Field()
	save.Master = partispants
	save.TurnGamerId = 0

	gameID, err := d.NewGame(save, password, gamerID, token)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Created new gameID: "+gameID+", by token:", token)
	w.Write([]byte("{\"game_id\":\"" + gameID + "\"}"))
}

type helperParse struct {
	Token    string `json:"token"`
	Password string `json:"password"`
	saveLoad.Participants
	gamerID int `json:"gamer_id"`
}

func Parse(i io.ReadCloser) (
	string,
	string,
	saveLoad.Participants,
	int,
	error,
) {
	data := make([]byte, 1024)
	n, err := i.Read(data)
	if err != nil && err != io.EOF {
		return "", "", saveLoad.Participants{}, 0, err
	}

	var helper helperParse
	err = json.Unmarshal(data[:n], &helper)
	if err != nil {
		return "", "", saveLoad.Participants{}, 0, err
	}
	return helper.Token, helper.Password, helper.Participants,
		helper.gamerID, nil
}
