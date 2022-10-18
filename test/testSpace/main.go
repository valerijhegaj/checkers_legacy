package main

import "fmt"

type helper struct {
	Password string `json:"password"`
	Name     string `json:"name"`
}

func main() {
	m := make(map[int]helper)
	m[1] = helper{"1", "1"}
	h := m[1]
	h.Password = "12"
	m[1] = h
	fmt.Println(m)
}
