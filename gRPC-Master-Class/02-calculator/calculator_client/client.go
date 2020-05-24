package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
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

func doServerStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeNumberDecompositionServer Streaming RPC...")
	req := &calcpb.PrimeNumberDecompositionRequest{
		Number: 504,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}

}

func doClientStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage: %v\n", err)
	}

	numbers := []float32{1, 2, 3, 4, 503}

	for _, number := range numbers {
		req := &calcpb.ComputeAverageRequest{
			Number: number,
		}
		log.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage: %v\n", err)
	}

	fmt.Printf("ComputeAverageResponse: %v\n", res.Average)
}
