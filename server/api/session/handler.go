package session

import (
	"io"
	"log"
	"net/http"

	"checkers/server/api"
	"checkers/server/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(
			"Bad method for new session, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to create new user:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedBody, err := api.Parse(body)
	username, password, maxAge :=
		parsedBody.UserName, parsedBody.Password, parsedBody.MaxAge
	if err != nil {
		log.Println("Failed new session: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storage := data.GetGlobalStorage()

	token, err := storage.NewSession(username, password)
	if err != nil {
		log.Println("Failed new session: " + err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cookie := &http.Cookie{Name: "token", Value: token, MaxAge: maxAge}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusCreated)
	log.Printf("Succesfuly new session: %s\n", username)
}
