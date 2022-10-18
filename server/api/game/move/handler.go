package move

import (
	"io"
	"log"
	"net/http"

	"checkers/server/api"
	"checkers/server/internal/data"
	"checkers/server/internal/errorsStrings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(
			"Bad method for move, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed move:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedBody, err := api.Parse(body)
	gameName, from, way :=
		parsedBody.GameName, parsedBody.From, parsedBody.Way
	if err != nil {
		log.Println("Failed move: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var token string
	cookies := r.Cookies()
	for _, c := range cookies {
		if c.Name == "token" {
			token = c.Value
		}
	}

	storage := data.GetGlobalStorage()
	err = storage.MakeMove(token, gameName, from, way)
	if err != nil {
		log.Println("Failed move: " + err.Error())
		switch err.Error() {
		case errorsStrings.NotAuthorized:
			w.WriteHeader(http.StatusForbidden)
		case errorsStrings.NotFound:
			w.WriteHeader(http.StatusNotFound)
		case errorsStrings.PermissionDenied:
			w.WriteHeader(http.StatusForbidden)
		case errorsStrings.IncorrectMove:
			w.WriteHeader(http.StatusMethodNotAllowed)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	log.Println("Successfully moved:", gameName)
	w.WriteHeader(http.StatusCreated)
}
