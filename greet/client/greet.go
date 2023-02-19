package main

import (
	"context"
	"log"

	pb "github.com/cmorales95/go-grpc-masterclass/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGret was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Cristian",
	})

	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
