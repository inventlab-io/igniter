package main

import (
	"github.com/igniter/command"
	"os"
)

func main() {
	//http.InitGin()

	os.Exit(command.Run(os.Args[1:]))
}
