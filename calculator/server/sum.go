package main

import (
	"context"
	"github.com/4925k/grpc-go/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, r *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("Received SumRequest %+v", r)

	return &proto.SumResponse{Sum: r.X + r.Y}, nil
}
