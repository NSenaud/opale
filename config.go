package opale

import (
	"github.com/BurntSushi/toml"
	"github.com/Sirupsen/logrus"
	"github.com/shibukawa/configdir"
)

const file = "config.toml"

var log = logrus.New()

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

func LoadConfig() (conf Config, err error) {
	configDirs := configdir.New("TechnicalDebtFactory", "opale")

	folder := configDirs.QueryFolderContainsFile(file)
	if folder != nil {
		// Deserialize configuration file.
		data, _ := folder.ReadFile(file)
		err = toml.Unmarshal(data, &conf)

		return
	} else {
		log.WithFields(logrus.Fields{
			"local":       configDirs.QueryFolders(configdir.Local)[0],
			"global":      configDirs.QueryFolders(configdir.Global)[0],
			"system":      configDirs.QueryFolders(configdir.System)[0],
			"config_file": file,
		}).Fatal("Can't find configiguration file.")
		panic("Can't load config.")
	}
}
