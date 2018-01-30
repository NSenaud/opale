package main

import (
	"log"
	"net"

	"github.com/NSenaud/opale"
	"github.com/NSenaud/opale/api"
	"google.golang.org/grpc"
)

func main() {
	// TODO Check params
	// ...

	// TODO Check config file
	// ...

	// Listener on Unix socket
	lis, err := net.Listen("unix", "/tmp/opale.sock")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach services to the server
	api.RegisterCpuServer(grpcServer, &s)
	api.RegisterRamServer(grpcServer, &s)

	// start the server in a goroutine
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// Monitor stuffs
	for {
		cpu, threads := opale.GetCpu()
		mem := opale.GetMem()

		opale.InsertIntoDb(cpu, threads, mem)
		// Sleep is not necessary yet since we already wait for a second in
		// GetCpu(), however it will be necessary as soon as we will be async,
		// and the interval will be setup in configuration file.
	}
}
