package main

import (
	"fmt"
	"log"
	"net"'
	pb 
	"google.golang.org/grpc"
)

const (
	port = "8080"
)

type helloServer struct {
	pb.RegisterGreetServiceServer
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatal("Error", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}