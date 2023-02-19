package main

import (
	"io"
	"log"

	pb "github.com/cmorales95/go-grpc-masterclass/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream %v\n", err)
		}

		res := "Hello " + req.FirstName + "!\n"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("error while sending data to client: %v\n", err)
		}
	}
}
