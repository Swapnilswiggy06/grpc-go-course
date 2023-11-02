package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	log.Printf("Greet response: %v", res.Result)
}
