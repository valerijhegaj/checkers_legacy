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

func main() {
	test_client_change_password("nam", "p", url)
}
