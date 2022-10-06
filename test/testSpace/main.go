package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type helper struct {
	Password string `json:"password"`
	Name     string `json:"name"`
}

func main() {
	s := "{\"name\":\"user\", \"password\":\"password\"}"
	var s1 helper
	err := json.Unmarshal([]byte(s), &s1)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(s1)
}
