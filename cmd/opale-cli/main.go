package main

import (
	"github.com/NSenaud/opale"
	"github.com/NSenaud/opale/client"
	log "github.com/sirupsen/logrus"
)

func main() {
	// TODO Check params
	// ...

	// Check config file
	// TODO Load config from XDG path
	config, err := opale.LoadConfig("config/example01.toml")
	if err != nil {
		// FIXME Should use default settings instead
		log.Fatal("Failed to read configuration file!")
	}

	// Load logger
	LogInit(&config)

	conn, c := client.IpcSubscribe(&config)
	defer conn.Close()

	for _, sensor := range config.Sensors {
		used_perc := client.GetUsedPercent(c, client.SensorEnum[sensor])
		log.Printf("%s: %.02f%s", sensor, used_perc, "%")
	}
}

func LogInit(conf *opale.Config) {
	if conf.Client.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
