package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	calcpb "github.com/ck3g/SomeDaysOfGo/gRPC-Master-Class/02-calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	summandOne := req.GetSum().GetSummandOne()
	summandTwo := req.GetSum().GetSummandTwo()
	result := summandOne + summandTwo
	res := &calcpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calcpb.PrimeNumberDecompositionRequest, stream calcpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)

	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calcpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Println("Divisor has increased to", divisor)
		}
	}

	return nil
}

func (*server) ComputeAverage(stream calcpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("ComputeAverage function was invoked with a streaming request")
	var sum float32
	var count int32
	var average float32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calcpb.ComputeAverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		number := req.GetNumber()
		sum += number
		count++
		average = sum / float32(count)

	}
}

func (*server) FindMax(stream calcpb.CalculatorService_FindMaxServer) error {
	fmt.Println("FindMax function was invoked with a streaming request")

	var maxNumber int64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
			return err
		}

		number := req.GetNumber()
		if number > maxNumber {
			maxNumber = number
		}

		err = stream.Send(&calcpb.FindMaxResponse{
			MaxNumber: maxNumber,
		})
		if err != nil {
			log.Fatalf("Error while sendind data to client: %v\n", err)
			return err
		}
	}
}

func main() {
	fmt.Println("The service is up and running...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
