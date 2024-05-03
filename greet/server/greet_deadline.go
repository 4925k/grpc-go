package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/4925k/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Println("GreetWithDeadline called")

	for i := 0; i < 5; i++ {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("deadline exceeded")
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		time.Sleep(1 * time.Second)
	}

	return &proto.GreetResponse{Result: "Hello" + req.FirstName}, nil
}
