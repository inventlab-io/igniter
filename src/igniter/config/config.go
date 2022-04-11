package config

type EtcdOptions struct {
	RequestTimeout    int
	Endpoints         []string
	ConnectionTimeout int
}

type StoreOptions struct {
	Type    string
	Options map[string]interface{}
}

type ServerConfig struct {
	Storage StoreOptions
}

type AgentConfig struct {
	Server struct {
		Address string `mapstructure:"ADDRESS"`
		Port    string `mapstructure:"PORT"`
	}
}
