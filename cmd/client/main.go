package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/stormcat24/skyray/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial(":5514", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewSkyrayServiceClient(conn)

	md := metadata.New(map[string]string{})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	req := pb.Command{
		Command: "./tick.sh",
	}

	stream, err := client.Connect(ctx, &req)
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("stream is closed")
				return
			}
			log.Printf("[ERROR] %s, ", err.Error())
		} else {
			fmt.Printf("# %s", string(res.Output))
		}
	}
}
