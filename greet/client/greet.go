package main

import (
	"context"
	"github.com/4925k/grpc-go/greet/proto"
	"log"
)

func doGreet(c proto.GreetServiceClient) {
	log.Println("start greet")

	res, err := c.Greet(context.Background(), &proto.GreetRequest{
		FirstName: "Dibek",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(res.Result)
}
