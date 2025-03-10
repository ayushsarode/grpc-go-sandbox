package main

import (
	"io"
	"log"

	pb "github.com/ayushsarode/grpc-go-sandbox/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Now we can use SendAndClose() because it's not a streaming response anymore
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}

		if err != nil {
			log.Printf("Error receiving stream: %v", err)
			return err
		}

		log.Printf("Received request with name: %v", req.Name)

		// Properly format the message
		messages = append(messages, "Hello "+req.Name)
	}
}
