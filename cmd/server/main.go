package main

import (
	"google.golang.org/grpc"
	pb "grpc-server/pkg/api"
	"grpc-server/pkg/user"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &user.GRPCServer{}
	pb.RegisterUserServer(s, srv)

	address := ":8080"
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Starting server at %s\n", address)

	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}
