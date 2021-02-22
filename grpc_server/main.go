package main

import (
	"grpc_server/geocoder"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {

	}
	s := geocoder.Server{}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	geocoder.RegisterBookServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(l); err != nil {

	}
}
