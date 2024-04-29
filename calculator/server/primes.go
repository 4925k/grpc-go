package main

import (
	"github.com/4925k/grpc-go/calculator/proto"
)

func (s *Server) Primes(r *proto.PrimesRequest, stream proto.CalculatorService_PrimesServer) error {
	n, k := int(r.X), 2

	for n > 1 {
		if n%k == 0 {
			err := stream.Send(&proto.PrimesResponse{Prime: int32(k)})
			if err != nil {
				return err
			}
			n = n / k
		} else {
			k++
		}
	}

	return nil
}
