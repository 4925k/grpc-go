package main

import (
	"context"
	"github.com/4925k/grpc-go/calculator/proto"
	"io"
	"log"
)

func doPrimes(c proto.CalculatorServiceClient) {
	log.Println("Starting primes...")

	req := proto.PrimesRequest{X: 120}

	stream, err := c.Primes(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while calling Primes RPC: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while calling Primes RPC: %v", err)
			return
		}

		log.Printf("recieved %d", resp.Prime)
	}

	return
}
