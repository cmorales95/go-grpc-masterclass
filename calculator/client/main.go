package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cmorales95/go-grpc-masterclass/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	// doSum(client)
	// doPrimes(client)
	// doAvg(client)
	// doMax(client)
	// doSqrt(client, 10)
	doSqrt(client, -2)
}
