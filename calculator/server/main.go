package main

import (
	"log"
	"net"

	pb "github.com/cmorales95/go-grpc-masterclass/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var addr = "0.0.0.0:50051"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on:%v\n", err)
	}

	var opts []grpc.ServerOption
	tls := true // change that to false if needed
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		credential, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("falied loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(credential))
	}

	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer(opts...)
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
