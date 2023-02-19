package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/cmorales95/go-grpc-masterclass/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	fmt.Println("doAvg was invoked")

	numbers := []int32{3, 5, 9, 54, 23}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("error while calling Avg: %v\n", err)
	}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)
		stream.Send(&pb.AvgRequest{Number: number})
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %.2f", res.Result)

}
