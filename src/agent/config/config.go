package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type AgentConfig struct {
	Server struct {
		Address string `mapstructure:"ADDRESS"`
		Port    string `mapstructure:"PORT"`
	}
}

func LoadAgentConfig(path string) (config AgentConfig, err error) {
	agentV := viper.New()

	agentV.SetDefault("Server.Address", "http://localhost")
	agentV.SetDefault("Server.Port", 5050)

	if path != "" {

		dir := filepath.Dir(path)
		if dir != "" {
			agentV.AddConfigPath(dir)
		}

		fn := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		if fn == "" {
			fn = "ignition"
		}
		agentV.SetConfigName(fn)

		err = agentV.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}

	err = agentV.Unmarshal(&config)

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return
}
