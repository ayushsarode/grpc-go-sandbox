package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ayushsarode/grpc-go-sandbox/proto"
)



func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("couldn't greet %v", err)
	}
	log.Printf("%s", res.Message)
}