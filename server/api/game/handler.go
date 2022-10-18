package game

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
	gameName, password := parsedBody.GameName, parsedBody.Password
	if err != nil {
		log.Println("Failed new game: " + err.Error())
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

	switch r.Method {
	case http.MethodGet:
		get(w, token, gameName)
	case http.MethodPost:
		login(w, token, gameName, password)
	case http.MethodPut:
		change(w, token, gameName, body)
	case http.MethodDelete:
		del(w, token, gameName)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func get(w http.ResponseWriter, token, gameName string) {
	storage := data.GetGlobalStorage()
	game, err := storage.GetGame(token, gameName)
	if err == nil {
		w.Write(game)
		w.WriteHeader(http.StatusOK)
		return
	}
	switch err.Error() { //refactor extract method
	case errorsStrings.NotAuthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case errorsStrings.NotFound:
		w.WriteHeader(http.StatusNotFound)
	case errorsStrings.PermissionDenied:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, token, gameName, password string) {
	storage := data.GetGlobalStorage()
	err := storage.LoginGame(token, gameName, password)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}
	switch err.Error() {
	case errorsStrings.NotAuthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case errorsStrings.NotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func change(
	w http.ResponseWriter, token, gameName string,
	body []byte,
) {
	var settings game.Settings
	err := json.Unmarshal(body, &settings)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storage := data.GetGlobalStorage()
	err = storage.ChangeGame(token, gameName, settings)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		return
	}
	switch err.Error() {
	case errorsStrings.NotAuthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case errorsStrings.NotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func del(w http.ResponseWriter, token, gameName string) {
	storage := data.GetGlobalStorage()
	err := storage.DeleteGame(token, gameName)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		return
	}
	switch err.Error() {
	case errorsStrings.NotAuthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case errorsStrings.NotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

}
