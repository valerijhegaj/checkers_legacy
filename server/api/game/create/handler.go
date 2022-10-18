package create

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"checkers/server/api"
	"checkers/server/internal/data"
	"checkers/server/internal/errorsStrings"
	"checkers/server/internal/game"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(
			"Bad method for new game, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed new game:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedBody, err := api.Parse(body)
	gameName, password :=
		parsedBody.GameName, parsedBody.Password
	if err != nil {
		log.Println("Failed new game: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	settings := game.Settings{}
	err = json.Unmarshal(body, &settings)
	if err != nil {
		log.Println("Failed new game: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storage := data.GetGlobalStorage()
	err = storage.NewGame(gameName, password, settings)
	if err != nil {
		log.Println("Failed new game: " + err.Error())
		if err.Error() == errorsStrings.GameAlreadyExist {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Successfully created game:", gameName)
}
