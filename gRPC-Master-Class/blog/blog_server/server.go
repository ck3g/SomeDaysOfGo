package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ck3g/SomeDaysOfGo/gRPC-Master-Class/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failied to listen: %v", err)
	}

	s := grpc.NewServer()

	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server...")

		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for Control+C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	fmt.Println("The service is up and running...")

	// Block until a signal is received
	<-ch
	fmt.Println()
	fmt.Println("Stopping the server...")
	s.Stop()
	fmt.Println("Closing the listener...")
	lis.Close()
	fmt.Println("End of program")
}
