package main

import (
	"fmt"
	"log"
	"net"

	"github.com/stormcat24/skyray/config"
	"github.com/stormcat24/skyray/endpoint"
	"github.com/stormcat24/skyray/pb"
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
	pb.RegisterSkyrayServiceServer(server, ep)

	log.Printf("skyray listening port %s\n", conf.Port)
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
