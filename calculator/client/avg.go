package main

import (
	"context"
	"github.com/4925k/grpc-go/calculator/proto"
	"log"
	"time"
)

func doAvg(client proto.CalculatorServiceClient) {
	nums := []int32{1, 2, 3, 4}

	stream, err := client.Avg(context.Background())
	if err != nil {
		log.Fatalf("%v.Avg(_) = _, %v", client, err)
	}

	for _, n := range nums {
		log.Println("send request for avg", n)

		err := stream.Send(&proto.AvgRequest{Num: n})
		if err != nil {
			log.Fatalf("%v.Avg(_) = _, %v", stream, err)
		}

		time.Sleep(1 * time.Second)
	}

	avg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.Avg(_) = _, %v", stream, err)
	}

	log.Println("avg", avg)
}
