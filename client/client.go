package client

import (
	"context"
	"net"
	"time"

	"github.com/NSenaud/opale"
	"github.com/NSenaud/opale/api"
	"github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

var log = logrus.New()

type sensorType int
type GetUsedPercentFunc func(api.OpaleClient, context.Context, *api.StatusRequest, ...grpc.CallOption) (*api.UsedPercent, error)

const (
	cpu sensorType = iota
	ram
	MaxSensorType
)

var SensorEnum = map[string]sensorType{
	"cpu": cpu,
	"ram": ram,
}

var getUsedPercentFuncs = map[sensorType]GetUsedPercentFunc{
	cpu: api.OpaleClient.GetCpuUsedPercent,
	ram: api.OpaleClient.GetRamUsedPercent,
}

func IpcSubscribe(conf *opale.Config) (conn *grpc.ClientConn, client api.OpaleClient) {
	conn, err := grpc.Dial(conf.Socket,
		grpc.WithInsecure(),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	client = api.NewOpaleClient(conn)

	return
}

func GetUsedPercent(client api.OpaleClient, sensor sensorType) float64 {
	log.WithFields(logrus.Fields{
		"client": client,
		"sensor": sensor,
	}).Debug("client.GetUsedPercent")

	log.Debug("Get func from sensor name...")
	if sensor < MaxSensorType {
		f, ok := getUsedPercentFuncs[sensor]
		if ok {
			res, err := f(client, context.Background(), &api.StatusRequest{})
			if err != nil {
				log.Fatal("Error while calling f:", f, "->", err)
			}
			return res.Value
		} else {
			log.Fatal("Error while mapping sensorType", sensor, "with getFuncs.")
		}
	}
	return 0
}
