package main

import (
	protos "grpc-service/generated-protos"
	"grpc-service/server"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {
	address := ":8080"
	grpcServer := grpc.NewServer()

	gameCatalogServer := server.New()

	protos.RegisterGameCatalogServiceServer(grpcServer, gameCatalogServer)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			grpcServer.GracefulStop()
			os.Exit(1)
		}
	}()

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
