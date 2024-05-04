package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	pb "github.com/4925k/grpc-go/greet/proto"
)

var (
	addr = "0.0.0.0:50051"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on %s", addr)

	tls := true
	opts := []grpc.ServerOption{}
	if tls {
		creds, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
