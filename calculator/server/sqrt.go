package main

import (
	"context"
	"github.com/4925k/grpc-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Server) Sqrt(_ context.Context, req *proto.SqrtRequest) (*proto.SqrtResponse, error) {
	log.Printf("Received Sqrt request with params %+v", req)

	// check for negative number
	number := req.Number
	if number <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Number must be positive")
	}

	return &proto.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
