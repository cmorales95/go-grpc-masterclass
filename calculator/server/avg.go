package main

import (
	"io"
	"log"

	pb "github.com/cmorales95/go-grpc-masterclass/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg functions was invoked")

	var count uint
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{Result: float64(sum) / float64(count)})
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving req: %d\n", req.Number)
		sum += req.Number
		count++
	}
}
