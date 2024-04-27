package main

import (
	"context"
	"github.com/4925k/grpc-go/calculator/proto"
	"log"
)

func doSum(client proto.CalculatorServiceClient) {
	log.Println("starting sum client")

	res, err := client.Sum(context.Background(), &proto.SumRequest{X: 10, Y: 3})
	if err != nil {
		log.Fatal("error while calling Sum RPC:", err)
	}

	log.Println(res.Sum)
}
