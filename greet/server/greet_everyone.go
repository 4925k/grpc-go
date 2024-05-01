package main

import (
	"fmt"
	"github.com/4925k/grpc-go/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream proto.GreetService_GreetEveryoneServer) error {
	log.Printf("Entering server.GreetEveryone")

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		res := fmt.Sprintf("Hello %s!", in.FirstName)
		err = stream.Send(&proto.GreetResponse{Result: res})
		if err != nil {
			return err
		}
	}
}
