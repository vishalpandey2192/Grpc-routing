package main

import (
	"Grpc-routing/calculator/calculatorpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am from Sum client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// fmt.Println("Created a client : %f", c)

	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting a Unary Calculator RPC")

	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 20,
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling the Sum RPC %v", err)
	}

	log.Printf("Response from Sum GRPC %v", res)
}
