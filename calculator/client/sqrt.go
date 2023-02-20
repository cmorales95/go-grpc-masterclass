package main

import (
	"context"
	"log"

	pb "github.com/cmorales95/go-grpc-masterclass/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("error message from server: %s\n", e.Message())
			log.Printf("error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("we probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("a non gRPC error: %v\n", err)
		}
	}
	log.Printf("Sqrt: %.2f", res.Result)
}
