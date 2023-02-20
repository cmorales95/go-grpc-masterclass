package main

import (
	"log"
	"net"

	pb "github.com/cmorales95/go-grpc-masterclass/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr = "0.0.0.0:50051"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on:%v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

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

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
