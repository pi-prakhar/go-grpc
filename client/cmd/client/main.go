package main

import (
	"log"

	pb "github.com/pi-prakhar/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NameList{
	// 	Names: []string{"Akhil", "Alice", "Bob"},
	// }

	callSayHello(client)
}
