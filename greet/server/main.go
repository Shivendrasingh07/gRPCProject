package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/Shivendrasingh07/gRPCProject/greet/proto"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen server on port %v\n", err)
	}
	log.Printf("listening on %v\n ", addr)

	s := grpc.NewServer()

	//srv := Server{}

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failled to serve : %v\n", err)
	}

}
