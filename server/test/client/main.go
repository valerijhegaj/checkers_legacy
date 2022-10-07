package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const url = "http://localhost:4444"

func test_client_get_token(user, password, url string) {
	client := &http.Client{}
	body := "{\"name\":\"" + user + "\", \"password\":\"" + password + "\"}"
	r, err := http.NewRequest(
		http.MethodPost,
		url+"/api/get_token",
		strings.NewReader(body),
	)
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func test_client_create_user(user, password, url string) {
	client := &http.Client{}
	body := "{\"name\":\"" + user + "\", \"password\":\"" + password + "\"}"
	r, err := http.NewRequest(
		http.MethodPost,
		url+"/api/create_user",
		strings.NewReader(body),
	)
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func test_client_change_password(token, password, url string) {
	client := &http.Client{}
	body := "{\"token\":\"" + token + "\", " +
		"\"password\":\"" + password + "\"}"
	r, err := http.NewRequest(
		http.MethodPatch,
		url+"/api/change_password",
		strings.NewReader(body),
	)
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func test_client_create_game(token, url string) {
	client := &http.Client{}
	body := "{\"token\": \"" + token + "\", \"password\": \"password\"," +
		"\"gamer0\": 0,\"level0\": 3, \"gamer1\": 1, \"level1\": 3," +
		"\"gamer_id\": 0 }"
	r, err := http.NewRequest(
		http.MethodPost,
		url+"/api/create_game",
		strings.NewReader(body),
	)
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func main() {
	test_client_create_user("nam", "p", url)
	test_client_create_game("nam", url)
	test_client_get_token("nam", "p", url)
	test_client_create_game("nam", url)

}
