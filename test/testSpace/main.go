package main

import (
	"fmt"

	"checkers/core"
)

type helper struct {
	Password string `json:"password"`
	Name     string `json:"name"`
}

func main() {
	checker_1 := core.Checker{1}
	checker_1_clone := core.Checker{1}
	checker_0 := core.Checker{0}
	//king_1 := core.King{1}
	//king_0 := core.King{0}
	fmt.Println(checker_1 == checker_1_clone)
	fmt.Println(checker_0 == checker_1)
	fmt.Println(core.Figure(checker_1_clone) == core.Figure(checker_1))
}
