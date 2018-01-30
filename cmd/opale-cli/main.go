package main

import (
	"log"
	"net"
	"time"

	"github.com/NSenaud/opale/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// TODO Check params
	// ...

	// TODO Check config file
	// ...

	// Open Unix socket
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("/tmp/opale.sock",
		grpc.WithInsecure(),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// Subcribe to Cpu and Ram services
	c := api.NewCpuClient(conn)
	r := api.NewRamClient(conn)

	cpu, err := c.GetCpuInfo(context.Background(), &api.StatusRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetCpuInfo: %s", err)
	}
	log.Printf("Response from server: %s", cpu.UsedPercent)

	ram, err := r.GetRamInfo(context.Background(), &api.StatusRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetRamInfo: %s", err)
	}
	log.Printf("Response from server: %s", ram.UsedPercent)
}
