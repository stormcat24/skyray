package main

import (
	"fmt"
	"net"

	"github.com/stormcat24/skyray/config"
	"github.com/stormcat24/skyray/endpoint"
	"google.golang.org/grpc"
)

func main() {

	conf := config.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	ep := endpoint.NewSkyrayEndpoint()
	endpoint.RegisterSkyrayServiceServer(server, ep)

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
