//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/users.proto
package main

import (
	pb "github.com/andresorav/proto-go/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

type server struct {
	pb.UnimplementedUsersServer
}

func (s *server) One(ctx context.Context, in *pb.GetUser) (*pb.UserFound, error) {
	log.Printf("Received: %v ", in.GetId())
	return &pb.UserFound{Username: "userId is " + in.GetId()}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUsersServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


