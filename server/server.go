package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/leslesnoa/grpcdemo01/pb"
	"google.golang.org/grpc"
)

// ---service----
type GreeterService struct {
}

func (s *GreeterService) SayHello(ctx context.Context, message *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println(message)
	log.Printf("Received: %v", message.Name)
	time.Sleep(3 * time.Second)
	return &pb.HelloReply{Message: "Hello " + message.Name}, nil
}

// -------------

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &GreeterService{})

	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
