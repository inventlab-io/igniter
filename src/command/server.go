package command

import (
	"flag"
	"fmt"
	"github.com/igniter/server"
	"github.com/igniter/server/config"
	"io"
)

type ServerCommand struct {
	ShutdownCh chan struct{}
	SighupCh   chan struct{}
	SigUSR2Ch  chan struct{}

	logOutput io.Writer
}

func (s ServerCommand) Help() string {
	//TODO implement me
	return "TODO"
}

func (a ServerCommand) Run(args []string) int {

	var configFile string
	set := flag.NewFlagSet("server", flag.ExitOnError)
	set.StringVar(&configFile, "config", "", "full path of configuration file")
	set.Parse(args)

	cfg, err := config.LoadServerConfig(configFile)
	if err != nil {
		panic(fmt.Errorf("error loading config"))
	}

	srv := server.Server{}
	defer srv.Shutdown()

	srv.Run(cfg)

	return 0
}

func (s ServerCommand) Synopsis() string {
	//TODO implement me
	return "Syn"
}
