package main

import pb "github.com/cmorales95/go-grpc-masterclass/greet/proto"

type Server struct {
	pb.UnimplementedGreetServiceServer
}
