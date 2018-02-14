package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/NSenaud/opale"
	"github.com/NSenaud/opale/api"
	"github.com/NSenaud/opale/db"
	"github.com/NSenaud/opale/sensors"
	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// TODO Check params
	// ...

	// Load config file
	config, err := opale.LoadConfig()
	if err != nil {
		// FIXME Should use default settings instead
		log.Fatal("Failed to read configuration file!")
	}

	// Load logger
	LogInit(&config)

	// Listener on Unix socket
	lis, err := net.Listen("unix", config.Socket)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Listen on program exit to remove socket file.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		var err = os.Remove(config.Socket)
		if err != nil {
			log.Fatal("Failed to delete socket file:", config.Socket)
			return
		}

		log.Println("Deleted socket file.")
		os.Exit(0)
	}()

	// create a server instance
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach services to the server
	api.RegisterOpaleServer(grpcServer, &s)

	// start the server in a goroutine
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// Monitor stuffs
	// TODO Split in func and goroutine.
	// TODO Monitor sensors listed in config file.
	// TODO Sleep (from config file).
	// TODO Cleanup db (from config file).
	for {
		cpu, threads := sensors.GetCpu()
		mem := sensors.GetRam()

		db.InsertIntoDb(cpu, threads, mem)
		// Sleep is not necessary yet since we already wait for a second in
		// GetCpu(), however it will be necessary as soon as we will be async,
		// and the interval will be setup in configuration file.
	}
}

func LogInit(conf *opale.Config) {
	if conf.Client.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
