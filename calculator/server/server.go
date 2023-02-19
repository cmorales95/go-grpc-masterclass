package main

import pb "github.com/cmorales95/go-grpc-masterclass/calculator/proto"

type Server struct {
	pb.UnimplementedCalculatorServiceServer
}
