package agent

import (
	"fmt"
	"github.com/igniter/agent/config"
)

func Run(config config.AgentConfig) {
	fmt.Println("Agent at %s:%s", config.Server.Address, config.Server.Port)
	fmt.Println("Agent Running")
}
