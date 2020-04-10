package main

import (
	"context"
	"log"

	"github.com/luindayk/grpc-ok/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)

	Hello(err, client)

}

func Hello(err error, client pb.HelloServiceClient) {
	request := &pb.HelloRequest{
		Name: "Willyan Luindayk",
	}

	res, err := client.Hello(context.Background(), request)

	if err != nil {
		log.Fatalf("Error during the execution: %v", err)
	}

	log.Println(res)
}
