package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "try-grpc/proto"

	"google.golang.org/grpc"
)

// server is used to implement greeting.GreetingServiceServer
type server struct {
	pb.UnimplementedGreetingServiceServer
}

// SayHello implements greeting.GreetingServiceServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// SayHelloStream implements greeting.GreetingServiceServer
func (s *server) SayHelloStream(in *pb.HelloRequest, stream pb.GreetingService_SayHelloStreamServer) error {
	log.Printf("Received stream request for: %v", in.GetName())

	for i := 0; i < 5; i++ {
		reply := &pb.HelloReply{
			Message: fmt.Sprintf("Hello %s - Message %d", in.GetName(), i+1),
		}
		if err := stream.Send(reply); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
