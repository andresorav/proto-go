package main

import (
	pb "github.com/andresorav/proto-go/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const address = "192.168.46.100:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUsersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.One(ctx, &pb.GetUser{Id: "complicatedUserId"})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	log.Printf("Greetings: %s", r.GetUsername())
}