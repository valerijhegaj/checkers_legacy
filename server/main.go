package main

import (
	"log"
	"net/http"

	"checkers/server/api/change_password"
	createGame "checkers/server/api/create_game"
	"checkers/server/api/create_user"
	deleteUser "checkers/server/api/delete_user"
	getGame "checkers/server/api/get_game"
	"checkers/server/api/get_token"
	logInGame "checkers/server/api/login_game"
	logOutGame "checkers/server/api/logout_game"
	"checkers/server/api/move"
	"checkers/server/internal/data"
)

func main() {
	err := data.InitStorage()
	if err != nil {
		log.Println("can't initiate storage")
	}

	http.HandleFunc("/api/create_user", createUser.Handler)
	http.HandleFunc("/api/get_token", getToken.Handler)
	http.HandleFunc("/api/change_password", changePassword.Handler)
	http.HandleFunc("/api/delete_user", deleteUser.Handler)
	http.HandleFunc("/api/create_game", createGame.Handler)
	http.HandleFunc("/api/get_game", getGame.Handler)
	http.HandleFunc("/api/login_game", logInGame.Handler)
	http.HandleFunc("/api/logout_game", logOutGame.Handler)
	http.HandleFunc("/api/move", move.Handler)

	err = http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
