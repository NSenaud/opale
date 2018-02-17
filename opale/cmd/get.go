package cmd

import (
	"fmt"

	"github.com/NSenaud/opale"
	"github.com/NSenaud/opale/client"
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func get(config *opale.Config, sensor string) {
	log.WithFields(logrus.Fields{
		"config": config,
		"sensor": sensor,
	}).Debug("Get command called.")

	log.Debug("Subscribing to IPC server...")
	conn, c := client.IpcSubscribe(config)
	defer conn.Close()

	// TODO Check sensor is in config.Sensors

	log.Debug("Requesting used percentage...")
	used_perc := client.GetUsedPercent(c, client.SensorEnum[sensor])
	// TODO Fix print format
	fmt.Printf("%.02f%s", used_perc, "%")
}
