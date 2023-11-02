package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Received Primes RPC: %v", in)
	number := in.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{Result: divisor})
			number = number / divisor
		} else {
			divisor++
			log.Printf("Divisor has increased to %v", divisor)
		}
	}
	return nil
}
