package main

import (
	"context"
	"log"
	"net"

	pb "github.com/leslesnoa/grpcdemo01/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ---service----
type GreeterService struct {
}

func (s *GreeterService) SayHello(ctx context.Context, message *pb.HelloRequest) (*pb.HelloReply, error) {
	// log.Println(message)
	log.Printf("Received: %v", message.Name)
	// time.Sleep(3 * time.Second)
	return &pb.HelloReply{Message: "Hello " + message.Name}, nil
}

// -------------

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// -------TLS認証処理を追加-------
	cred, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(cred))
	// -----------------------------
	// s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &GreeterService{})

	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
