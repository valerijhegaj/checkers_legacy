package game

import (
	"log"
	"net/http"

	"checkers/server/api"
	"checkers/server/internal/data"
	"checkers/server/internal/errorsStrings"
	"checkers/server/pkg/defines"
	"checkers/server/pkg/file"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	api.EachHandlerRoutine(w)
	if r.Method == http.MethodOptions {
		api.CreateResponseCROPS(w, "GET, POST, PUT, DELETE")
		return
	}
	var token string
	cookies := r.Cookies()
	for _, c := range cookies {
		if c.Name == "token" {
			token = c.Value
		}
	}

	if r.Method == http.MethodGet {
		gameName := r.URL.Query().Get("gamename")
		get(w, token, gameName)
		return
	}
	body, err := file.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed new game:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedBody, err := api.Parse(body)
	password, settings, gameName :=
		parsedBody.Password, parsedBody.Settings, parsedBody.GameName
	if err != nil {
		log.Println("Failed new game: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		get(w, token, gameName)
	case http.MethodPost:
		login(w, token, gameName, password)
	case http.MethodPut:
		change(w, token, gameName, settings)
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
		log.Printf(
			"Successfully get game token: %s, gamename %s\n", token,
			gameName,
		)
		return
	}

	log.Printf(
		"Failed get game error: %s, token: %s, gamename: %s", err.Error(),
		token, gameName,
	)
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
		log.Printf(
			"Successfully log in game token: %s, gamename %s\n", token,
			gameName,
		)
		return
	}

	log.Printf(
		"Failed log in game error: %s, token: %s, gamename: %s",
		err.Error(),
		token, gameName,
	)
	switch err.Error() {
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

func change(
	w http.ResponseWriter, token, gameName string,
	settings defines.Settings,
) {
	storage := data.GetGlobalStorage()
	err := storage.ChangeGame(token, gameName, settings)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Printf(
			"Successfully change settings game token: %s, gamename %s\n",
			token,
			gameName,
		)
		return
	}

	log.Printf(
		"Failed change settings game error: %s, token: %s, gamename: %s",
		err.Error(),
		token, gameName,
	)
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
		log.Printf(
			"Successfully del game token: %s, gamename %s\n", token,
			gameName,
		)
		return
	}

	log.Printf(
		"Failed del game error: %s, token: %s, gamename: %s", err.Error(),
		token, gameName,
	)
	switch err.Error() {
	case errorsStrings.NotAuthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case errorsStrings.NotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

}
