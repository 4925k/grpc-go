package main

import (
	"context"
	"github.com/4925k/grpc-go/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(client proto.CalculatorServiceClient) {
	nums := []int32{1, 2, 4, 12, 13, 5, 6, 7, 8, 9, 10, 11, 14, 15, 16}

	waitc := make(chan struct{})

	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for _, num := range nums {
			log.Printf("Sending num %d", num)
			req := &proto.MaxRequest{X: num}
			stream.Send(req)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
			}

			if err != nil {
				log.Fatalln(err)
			}

			log.Printf("current Max %v", res.Max)
		}
	}()

	<-waitc
}
