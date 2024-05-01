package main

import (
	"context"
	"github.com/4925k/grpc-go/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c proto.GreetServiceClient) {
	log.Println("calling greetEveryone...")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while calling greetEveryone: %+v", err)
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "Dbk"},
		{FirstName: "Btsl"},
		{FirstName: "Nshn"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending request: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while receiving response: %+v", err)
				break
			}

			log.Printf("received response: %v", res)
		}

		close(waitc)
	}()

	<-waitc
}
