package main

import (
	"context"
	"github.com/4925k/grpc-go/greet/proto"
	"log"
	"time"
)

func doLongGreet(client proto.GreetServiceClient) {
	log.Println("start to do long greet")

	greetRequest := []*proto.GreetRequest{
		{FirstName: "Dbk"},
		{FirstName: "Nshn"},
		{FirstName: "Btsl"},
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("%v.LongGreet(_) = _, %v", client, err)
	}
	for _, req := range greetRequest {
		log.Printf("start to send req: %v\n", req)

		err := stream.Send(req)
		if err != nil {
			log.Fatalf("%v.LongGreet(_) = _, %v", stream, err)
		}

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseSend() got error %v, want %v", stream, err, nil)
	}

	log.Println(res)
}
