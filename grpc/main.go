package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "grpc/proto"

	"google.golang.org/grpc"

	ser "grpc/server"
)

func main() {

	go StartServer()
	startClient()

	time.Sleep(2 * time.Second)

}

func StartServer() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreetingServer(s, &ser.Server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startClient() {
	address := "localhost:50051"

	con, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer con.Close()

	c := pb.NewGreetingClient(con)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Rahul Chaube"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())
}
