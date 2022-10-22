package main

import (
	"log"
	"net/http"

	"checkers/server/api/game"
	"checkers/server/api/game/create"
	"checkers/server/api/game/move"
	"checkers/server/api/session"
	"checkers/server/api/user"
	"checkers/server/internal/data"
)

func main() {
	err := data.InitGlobalStorage()
	if err != nil {
		log.Println(
			"Storage initialize " +
				"finished with error: " + err.Error(),
		)
	}

	const PORT = ":4444"
	//http.HandleFunc("endpoint", handler)

	http.HandleFunc("/api/user", user.Handler)
	http.HandleFunc("/api/session", session.Handler)
	http.HandleFunc("/api/game/create", create.Handler)
	http.HandleFunc("/api/game/move", move.Handler)
	http.HandleFunc("/api/game", game.Handler)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err.Error())
	}
}
