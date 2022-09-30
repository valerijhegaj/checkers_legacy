package main

import "fmt"

func main() {
	var x [4]func()
	s := []string{"0", "1", "2", "3"}
	for i := 0; i < 4; i++ {
		c := s[i]
		x[i] = func() {

			fmt.Println(c)
		}
	}

	for _, f := range x {
		f()
	}

}
