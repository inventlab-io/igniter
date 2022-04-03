package agent

import (
	"fmt"
	"github.com/igniter/config"
)

func Run(cfg config.AgentConfig) {
	fmt.Println("Agent at %s:%s", cfg.Server.Address, cfg.Server.Port)
	fmt.Println("Agent Running")
}
