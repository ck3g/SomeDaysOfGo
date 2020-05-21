package main

import (
	"fmt"
	"log"

	greetpb "github.com/ck3g/SomeDaysOfGo/gRPC-Master-Class/01-greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Don't use SSL
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)
}
