package opale_test

import (
	"testing"

	"github.com/NSenaud/opale"
)

func TestNewSensorsFromConfig(t *testing.T) {
	conf, _ := opale.LoadConfig("config/example01.toml")
	sensors := opale.NewSensorsFromConfig(&conf)

	if !sensors.Cpu.Enabled || !sensors.Ram.Enabled {
		t.Error("Failed to properly load sensors from configuration file.")
	}
}
