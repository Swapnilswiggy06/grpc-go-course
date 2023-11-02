package main

import (
	"context"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func (*Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
