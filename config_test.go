package opale_test

import (
	"reflect"
	"testing"

	"github.com/NSenaud/opale"
)

func TestLoadFile(t *testing.T) {
	file, err := opale.LoadConfig("config/example01.toml")
	if err != nil {
		t.Error("Failed to load config/example01.toml configuration file.")
	}

	common := opale.CommonConfig{
		Socket:  "/tmp/opale.sock",
		Monitor: []string{"cpu", "ram"},
	}

	server := opale.ServerConfig{
		Debug:     true,
		Interval:  10,
		Retention: [3]uint32{1, 15, 365},
	}

	client := opale.ClientConfig{
		Debug: true,
	}

	conf := opale.Config{
		Common: common,
		Server: server,
		Client: client,
	}

	if res := reflect.DeepEqual(conf, file); res == false {
		t.Error("Failed to read configuration.")
	}
}
