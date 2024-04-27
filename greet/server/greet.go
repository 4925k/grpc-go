package main

import (
	"context"
	"github.com/4925k/grpc-go/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, r *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("Received GreetRequest %+v", r)

	return &proto.GreetResponse{Result: "Hello " + r.FirstName}, nil
}
