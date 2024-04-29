package main

import (
	"context"
	"github.com/4925k/grpc-go/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c proto.GreetServiceClient) {
	log.Println("doing greet many times")

	req := &proto.GreetRequest{FirstName: "Dibek"}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("%v.GreetManyTimes(_) = _, %v", c, err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GreetManyTimes(_) reading stream = _, %v", c, err)
		}

		log.Println("Greet many times done", res.Result)
	}

}
