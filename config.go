package opale

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Socket  string
	Sensors []string
	Server  ServerConfig
	Client  ClientConfig
}

type ServerConfig struct {
	Debug     bool
	Interval  uint32
	Retention [3]uint32
}

type ClientConfig struct {
	Debug bool
}

func LoadConfig(file_path string) (Config, error) {
	var conf Config
	_, err := toml.DecodeFile(file_path, &conf)

	return conf, err
}
