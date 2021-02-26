package main

import (
	"context"
	"flag"
	v1 "go-grpc-multiply/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {

	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewMultiplyServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Multiply
	req := v1.NumbersRequest{
		Api:     apiVersion,
		Number1: 5,
		Number2: 2,
	}

	res, err := c.MultiplyNumbers(ctx, &req)
	if err != nil {
		log.Fatalf("Multiply failed: %v", err)
	}

	log.Printf("Multyply result: <%+v>\n\n", res)
}
