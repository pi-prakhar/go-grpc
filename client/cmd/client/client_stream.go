package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pi-prakhar/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("send the request with name : %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("CLient streaming finished")

	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}

	log.Printf("%v", res.Messages)
}
