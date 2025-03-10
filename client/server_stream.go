package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ayushsarode/grpc-go-sandbox/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming Started ")

	stream, err :=  client.SayHelloServerStreaming(context.Background(), names)

	if err != nil{
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}

		log.Println(message)
	}

	log.Printf("streaming finished")
}