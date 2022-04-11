package command

import (
	"flag"
	"fmt"
	"github.com/igniter/agent"
	"github.com/igniter/config"
	"io"
	"log"
)

type AgentCommand struct {
	ShutdownCh chan struct{}
	SighupCh   chan struct{}

	logWriter io.Writer
	logger    log.Logger

	flagConfigs []string
}

func (a AgentCommand) Help() string {
	return "TODO"
}

func (a AgentCommand) Run(args []string) int {

	var configFile string
	set := flag.NewFlagSet("agent", flag.ExitOnError)
	set.StringVar(&configFile, "config", "", "full path of configuration file")
	set.Parse(args)

	cfg, err := config.LoadAgentConfig(configFile)
	if err != nil {
		panic(fmt.Errorf("error loading config"))
	}

	agent.Run(cfg)

	return 0
}

func (a AgentCommand) Synopsis() string {
	return "Start Igniter Agent"
}
