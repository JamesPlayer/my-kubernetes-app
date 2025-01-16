package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/JamesPlayer/my-kubernetes-app/microservice/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPingPongServiceServer
}

func (s *server) Ping(_ context.Context, in *pb.PingPongRequest) (*pb.PingPongReply, error) {
	log.Printf("Received: %v", in.GetMsg())
	return &pb.PingPongReply{
		Msg: "Pinged Microservice",
		Env: map[string]string{
			"MY_NODE_NAME": os.Getenv("MY_NODE_NAME"),
			"MY_POD_NAME":  os.Getenv("MY_POD_NAME"),
			"MY_POD_IP":    os.Getenv("MY_POD_IP"),
			"MY_SECRET":    os.Getenv("MY_SECRET"),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPingPongServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
