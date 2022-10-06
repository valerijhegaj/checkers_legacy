package getToken

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"server/internal/data"
	"server/internal/helper"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(
			"bad method for post new token, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name, password, err := Parse(r.Body)
	if err != nil || name == "" {
		log.Println("name: ", name, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := generate_token(name, password)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(helper.StringToHttpStatus(token))
		return
	}

	err = save_token(token, name)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(token))
	log.Println("Created token:", token+",", "for user:", name)
}

func generate_token(name, password string) (
	string,
	error,
) {
	d, err := data.GetStorage()
	if err != nil {
		return helper.HttpStatusToString(http.StatusInternalServerError), err
	}
	err = d.CheckAccess(name, password)
	if err != nil {
		return helper.HttpStatusToString(http.StatusForbidden),
			errors.New(
				"Checked access, result: " + err.Error() + ", " +
					"name: " + name + ", password: " + password,
			)
	}
	return name, nil
}

func save_token(token, name string) error {
	d, err := data.GetStorage()
	if err != nil {
		return err
	}
	return d.NewToken(token, name)
}

type helperParse struct {
	Password string `json:"password"`
	Name     string `json:"name"`
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
	return helper.Name, helper.Password, nil
}
