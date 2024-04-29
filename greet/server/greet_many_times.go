package main

import (
	"fmt"
	"github.com/4925k/grpc-go/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(r *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	log.Println("GreetManyTimes invoked")

	for i := 0; i < 3; i++ {
		res := fmt.Sprintf("Hello %s, %d", r.FirstName, i)

		err := stream.Send(&proto.GreetResponse{Result: res})
		if err != nil {
			return err
		}
	}

	return nil
}
