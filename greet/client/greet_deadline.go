package main

import (
	"context"
	"fmt"
	"github.com/4925k/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetDeadline(client proto.GreetServiceClient, timeout time.Duration) {
	fmt.Println("start to greet with deadline")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := proto.GreetRequest{FirstName: "DBK"}

	res, err := client.GreetWithDeadline(ctx, &req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Fatalf("timeout exceeded")
			}

			log.Fatalf("error: %v", e)
		} else {
			log.Fatalf("greet with deadline unexpected error: %v", err)
		}
		return
	}

	fmt.Println("Greet With Deadline: ", res.Result)
}
