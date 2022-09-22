package main

import (
	"chekers/core"
	_interface "chekers/interface"
)

func main() {
	var Interface _interface.Interface
	var Core core.GameCore
	exiter := make(chan int)
	Interface.Init(exiter, Core)
	<-exiter
}
