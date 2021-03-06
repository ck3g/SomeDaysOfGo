package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	calcpb "github.com/ck3g/SomeDaysOfGo/gRPC-Master-Class/02-calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// doClientStreaming(c)
	// doBiDiStreaming(c)
	doErrorUnary(c)
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

func doBiDiStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	waitc := make(chan struct{})

	numbers := []int64{1, 129, 28, 238, 503, 20}

	go func() {
		for _, number := range numbers {
			req := &calcpb.FindMaxRequest{
				Number: number,
			}
			fmt.Printf("Sending a number: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			fmt.Printf("Max number is %v\n", res.MaxNumber)
		}

		close(waitc)
	}()

	<-waitc
}

func doErrorUnary(c calcpb.CalculatorServiceClient) {

	// correct call
	doErrorCall(c, 10)

	// errror call
	doErrorCall(c, -10)
}

func doErrorCall(c calcpb.CalculatorServiceClient, number int32) {
	res, err := c.SquareRoot(context.Background(), &calcpb.SquareRootRequest{Number: number})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// Actual error from gRPC (user error)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
			}
		} else {
			// some other error
			log.Fatalf("Error calling SquareRoot: %v\n", err)
		}
		return
	}
	fmt.Printf("Result of square root of %v is %v\n", number, res.GetNumberRoot())
}
