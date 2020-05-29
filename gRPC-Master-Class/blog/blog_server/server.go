package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ck3g/SomeDaysOfGo/gRPC-Master-Class/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("The service is up and running...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failied to listen: %v", err)
	}

	s := grpc.NewServer()

	blogpb.RegisterBlogServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
