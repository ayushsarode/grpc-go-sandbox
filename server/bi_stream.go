package main

import (
	"io"
	"log"
	pb "github.com/ayushsarode/grpc-go-sandbox/proto"
)

func (s *helloServer) callSayHelloBidirectionalStream(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil 

		}

		if err != nil {
			return err
		}

		log.Printf("Got req with name: %v", req.Name)
	

	res := &pb.HelloResponse{
		Message: "Hello" + req.Name,
	}
	if err := stream.Send(res); err != nil {
		return err
	}
	}
}