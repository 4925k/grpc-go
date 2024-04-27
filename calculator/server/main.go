package main

import (
	"github.com/4925k/grpc-go/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	addr = "0.0.0.0:50052"
)

type Server struct {
	proto.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
