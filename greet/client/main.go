package main

import (
	pb "github.com/Shivendrasingh07/gRPCProject/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to the server  %v\n", err)
	}
	log.Printf("connected to the server successfully")
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
}
