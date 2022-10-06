package changePassword

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"server/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		log.Println(
			"Bad method for change password, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, password, err := Parse(r.Body)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := data.GetStorage()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = d.ChangePassword(token, password)
	if err != nil {
		log.Println(
			"Tried to change password token:", token+", but",
			err.Error(),
		)
		if err.Error() == data.ErrorBadToken {
			w.WriteHeader(http.StatusForbidden)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	log.Println("Changed password for token:", token)
}

type helperParse struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func Parse(i io.ReadCloser) (
	string,
	string,
	error,
) {
	data := make([]byte, 1024)
	n, err := i.Read(data)
	if err != nil && err != io.EOF {
		return "", "", err
	}

	var helper helperParse
	err = json.Unmarshal(data[:n], &helper)
	if err != nil {
		return "", "", err
	}
	return helper.Token, helper.Password, nil
}
