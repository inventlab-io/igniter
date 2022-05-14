package main

import (
	"github.com/orikami/command"
	"os"
)

func main() {
	//http.InitGin()

	os.Exit(command.Run(os.Args[1:]))
}
