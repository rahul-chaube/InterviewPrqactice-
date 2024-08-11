package server

import (
	"context"
	"fmt"
	pb "grpc/proto"
)

type Server struct {
	pb.UnimplementedGreetingServer
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	fmt.Println("Server is called ")
	return &pb.HelloResponse{Name: "Hello"}, nil
}
