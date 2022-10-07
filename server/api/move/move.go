package move

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"checkers/core"
	"checkers/server/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		log.Println(
			"bad method for move, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, gameID, from, way, err := Parse(r.Body)
	if err != nil {
		log.Println("Tried to move, but " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := data.GetStorage()
	if err != nil {
		log.Println("Tried to move, but " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = d.Move(token, gameID, from, way)
	if err != nil {
		log.Println("Tried to move, but " + err.Error())
		if err.Error() == data.ErrorNotFoundGame {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == data.ErrorNotHaveAccess {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	log.Println("Moved gameID: " + gameID + ", token: " + token)
}

type helperParse struct {
	Token  string            `json:"token"`
	GameID string            `json:"game_id"`
	From   core.Coordinate   `json:"from"`
	Way    []core.Coordinate `json:"way"`
}

func Parse(i io.ReadCloser) (
	string,
	string,
	core.Coordinate,
	[]core.Coordinate,
	error,
) {
	data := make([]byte, 1024)
	n, err := i.Read(data)
	if err != nil && err != io.EOF {
		return "", "", core.Coordinate{}, nil, err
	}

	var helper helperParse
	err = json.Unmarshal(data[:n], &helper)
	if err != nil {
		return "", "", core.Coordinate{}, nil, err
	}
	return helper.Token, helper.GameID, helper.From, helper.Way, nil
}
