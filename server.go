package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	poker "grpc-demo/poker"
	pb "grpc-demo/proto"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedPokerServer
}

func (s *server) GetNuts(ctx context.Context, req *pb.GetNutsRequest) (*pb.GetNutsResponse, error) {
	res, err := poker.PokerEvaluator(req.Hand, req.River)
	return &pb.GetNutsResponse{
		Card: res,
	}, err
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPokerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
