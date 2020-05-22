package main

import (
	"context"
	"fmt"
	"log"

	calcpb "github.com/ck3g/SomeDaysOfGo/gRPC-Master-Class/02-calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := calcpb.NewCalculatorServiceClient(conn)

	doUnary(c)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	req := &calcpb.SumRequest{
		Sum: &calcpb.Sum{
			SummandOne: 1,
			SummandTwo: 2,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v\n", err)
	}

	log.Printf("Response from Sum: 1 + 2 = %v", res.Result)
}
