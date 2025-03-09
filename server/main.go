package main

import (
	"log"
	"net"
	pb "github.com/ayushsarode/grpc-go-sandbox/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main(){
	//listening to tcp port
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatal("Error", err)
	}

	// gRPC server which has no service registered and has not started to accept requests yet.
	grpcServer := grpc.NewServer()


	// Register the GreetService with the gRPC server
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}