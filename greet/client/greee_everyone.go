package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/cmorales95/go-grpc-masterclass/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Cristian"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request :%v\n", req.FirstName)
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("error while receiving: %v\n", err)
			}

			log.Printf("received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
