package deleteUser

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"checkers/server/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		log.Println(
			"Bad method for delete user, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := Parse(r.Body)
	if err != nil {
		log.Println("Tried to delete user, but " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := data.GetStorage()
	if err != nil {
		log.Println("Tried to delete user, but " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	d.DeleteUser(token)

	log.Println("Deleted: " + token)
}

type helperParse struct {
	Token string `json:"token"`
}

func Parse(i io.ReadCloser) (
	string,
	error,
) {
	data := make([]byte, 1024)
	n, err := i.Read(data)
	if err != nil && err != io.EOF {
		return "", err
	}

	var helper helperParse
	err = json.Unmarshal(data[:n], &helper)
	if err != nil {
		return "", err
	}
	return helper.Token, nil
}
