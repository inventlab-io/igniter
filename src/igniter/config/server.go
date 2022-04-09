package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func LoadServerConfig(path string) (cfg ServerConfig, err error) {
	agentV := viper.New()

	agentV.SetDefault("Storage.Type", "etcd")
	agentV.SetDefault("Storage.Options.RequestTimeout", 2)
	agentV.SetDefault("Storage.Options.Endpoints", []string{"127.0.0.1:2379"})
	agentV.SetDefault("Storage.Options.ConnectionTimeout", 2)

	agentV.SetConfigType("yaml")

	if path != "" {

		dir := filepath.Dir(path)
		if dir != "" {
			agentV.AddConfigPath(dir)
		}

		fn := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		if fn == "" {
			fn = "igniter"
		}

		agentV.SetConfigName(fn)

		err = agentV.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}

	err = agentV.Unmarshal(&cfg)

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return
}
