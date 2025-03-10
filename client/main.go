package main

import (
	"log"
	pb "github.com/ayushsarode/grpc-go-sandbox/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	//dial to tcp 8080 port
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("didnt connect %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)


	names := &pb.NameList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	// callSayHello(client)


	// callSayHelloServerStream(client, names)

	// callSayHelloClientStream(client, names)

	callSayHelloBidirectionalStream(client, names)
}