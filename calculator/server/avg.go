package main

import (
	"github.com/4925k/grpc-go/calculator/proto"
	"io"
	"log"
	"time"
)

func (s *Server) Avg(stream proto.CalculatorService_AvgServer) error {
	log.Printf("avg request: %v", time.Now())

	var avg int32
	var n int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.AvgResponse{Avg: float64(avg) / float64(n)})
		}
		if err != nil {
			log.Fatalf("Error while receiving request: %v", err)
		}

		log.Printf("Got request: %v", req.Num)

		time.Sleep(2 * time.Second)
		avg += req.Num
		n++
	}

	return nil
}
