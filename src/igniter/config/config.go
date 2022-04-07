package config

type EtcdOptions struct {
	RequestTimeout    int
	Endpoints         []string
	ConnectionTimeout int
}

type StoreOptions struct {
	StorageType string
}

type ServerConfig struct {
	RequestTimeout int
	Storage        string
	Etcd           EtcdOptions
}

type AgentConfig struct {
	Server struct {
		Address string `mapstructure:"ADDRESS"`
		Port    string `mapstructure:"PORT"`
	}
}
