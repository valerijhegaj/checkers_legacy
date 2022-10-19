package user

import (
	"log"
	"net/http"

	"checkers/server/api"
	"checkers/server/internal/data"
	"checkers/server/pkg/file"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		post(w, r)
	case http.MethodGet:
		get(w, r)
	default:
		log.Println(
			"Bad method for new user, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	body, err := file.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to create new user:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedBody, err := api.Parse(body)

	if err != nil {
		log.Println("Failed to create new user: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username, password := parsedBody.UserName, parsedBody.Password
	if username == "" {
		log.Println("Failed to create new user: no user name")
	}
	storage := data.GetGlobalStorage()
	err = storage.NewUser(username, password)
	if err != nil {
		log.Println("Failed to create new user: " + err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Successfully new user: " + username)
}

func get(w http.ResponseWriter, r *http.Request) {
	var token string
	cookies := r.Cookies()
	for _, c := range cookies {
		if c.Name == "token" {
			token = c.Value
		}
	}

	storage := data.GetGlobalStorage()
	_, err := storage.GetUserID(token)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
