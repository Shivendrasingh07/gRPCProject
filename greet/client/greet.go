package main

import (
	"context"
	pb "github.com/Shivendrasingh07/gRPCProject/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet function is invoked ")
	resp, err := client.Greet(context.Background(), &pb.GreetRequest{
		Message: " Hi there! Server great to have you",
	})
	if err != nil {
		log.Fatalf("failed to send request to the server %v\n", err)
	}
	log.Printf("Greeting from server : %v\n ", resp.Result)
}
