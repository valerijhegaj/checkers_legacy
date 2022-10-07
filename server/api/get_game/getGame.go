package getGame

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"checkers/server/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println(
			"bad method for get game, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, gameID, password, err := Parse(r.Body)
	if err != nil {
		log.Println("Tried to get game, but " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := data.GetStorage()
	if err != nil {
		log.Println("Tried to get game, but " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	save, err := d.GetGame(token, gameID, password)
	if err != nil {
		log.Println("Tried to get game, but " + err.Error())
		if err.Error() == data.ErrorNotFoundGame {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == data.ErrorNotHaveAccess {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	rawSave, err := save.GetRawSave()
	if err != nil {
		log.Println("Tried to get game, but " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Get game gameID: " + gameID + ", token: " + token)
	w.Write(rawSave)
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
