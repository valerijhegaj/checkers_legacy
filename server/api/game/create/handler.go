package create

import (
	"log"
	"net/http"

	"checkers/server/api"
	"checkers/server/internal/data"
	"checkers/server/internal/errorsStrings"
	"checkers/server/pkg/file"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	api.EachHandlerRoutine(w)

	if r.Method == http.MethodOptions {
		api.CreateResponseCROPS(w, "POST")
		return
	}

	if r.Method != http.MethodPost {
		log.Println(
			"Bad method for new game, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := file.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed new game:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedBody, err := api.Parse(body)
	gameName, password, settings :=
		parsedBody.GameName, parsedBody.Password, parsedBody.Settings
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
