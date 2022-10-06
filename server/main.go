package main

import (
	"log"
	"net/http"

	changePassword "server/api/change_password"
	createUser "server/api/create_user"
	"server/api/get_token"
	"server/internal/data"
)

func main() {
	err := data.InitStorage()
	if err != nil {
		log.Println("can't initiate storage")
	}

	http.HandleFunc("/api/get_token", getToken.Handler)
	http.HandleFunc("/api/create_user", createUser.Handler)
	http.HandleFunc("/api/change_password", changePassword.Handler)

	err = http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
