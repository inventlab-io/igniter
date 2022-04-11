package command

import (
	"fmt"
	"github.com/mitchellh/cli"
)

func Run(args []string) int {
	initCommands()
	return setup(args)
}

func setup(args []string) int {

	cli := &cli.CLI{
		Name:                       "ignition",
		Args:                       args,
		Commands:                   Commands,
		Autocomplete:               true,
		AutocompleteNoDefaultFlags: true,
	}

	exitCode, err := cli.Run()

	if err != nil {
		fmt.Errorf("Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
