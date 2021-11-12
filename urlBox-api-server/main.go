package main

import (
	"context"
	"log"

	"net"

	pb "ekstrah.com/go-protoBox-grpc"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type GenURLManagementServer struct {
	pb.UnimplementedGenURLManagementServer
}

// GenNewURL(context.Context, *ExURLReq) (*ExURLRes, error)
func (s *GenURLManagementServer) GenNewURL(ctx context.Context, req *pb.ExURLReq) (*pb.ExURLRes, error) {
	// string oriURL = 1;
	// string newURL = 2;
	// string userID = 3;
	// int32 count = 4;
	return &pb.ExURLRes{OriURL: "https://google.com", NewURL: "https://naver.com", UserID: "ekstrah", Count: 4}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Server Failed to listen to port %v", port)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGenURLManagementServer(grpcServer, &GenURLManagementServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Server failed to start due to %v", err)
	}
}
