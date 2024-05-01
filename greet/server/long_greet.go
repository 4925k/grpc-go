package main

import (
	"fmt"
	"github.com/4925k/grpc-go/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	log.Println("long greet called with stream")

	var res string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{Result: res})
		}
		if err != nil {
			log.Fatalf("error receiving from client %v", err)
		}

		res += fmt.Sprintf("Hello %s\n", req.FirstName)
	}
}
