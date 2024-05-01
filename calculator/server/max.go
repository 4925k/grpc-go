package main

import (
	"github.com/4925k/grpc-go/calculator/proto"
	"io"
)

func (s *Server) Max(stream proto.CalculatorService_MaxServer) error {
	var curMax int32

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		curMax = max(curMax, in.X)
		err = stream.Send(&proto.MaxResponse{Max: curMax})
		if err != nil {
			return err
		}
	}
}
