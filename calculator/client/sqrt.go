package main

import (
	"context"
	"github.com/4925k/grpc-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(client proto.CalculatorServiceClient) {
	log.Println("calling sqrt rpc")

	res, err := client.Sqrt(context.Background(), &proto.SqrtRequest{Number: -25})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("grpc call failed with code %s", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Fatalf("invalid argument - negative number")
			}
		}

		log.Fatalf("error: %v", e)
	}

	log.Printf("result: %v", res.Result)
}
