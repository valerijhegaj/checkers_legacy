package createUser

import (
	"log"
	"net/http"

	getToken "server/api/get_token"
	"server/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(
			"bad method for post new user, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name, password, err := getToken.Parse(r.Body)
	if err != nil || name == "" {
		log.Println("name: ", name, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := data.GetStorage()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = d.NewUser(name, password)
	if err != nil {
		log.Println("Tryed to create user:", name+", but", err.Error())
		if err.Error() == "already exist" {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	log.Println("Created new user:", name)
}
