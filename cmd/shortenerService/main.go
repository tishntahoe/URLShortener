package main

import (
	"github.com/tishntahoe/UrlShortener/internal/shotener"
	pb "github.com/tishntahoe/UrlShortener/proto/shortenerpb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		// logger
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterShortenerServiceServer(grpcServer, &shotener.Server{})

	if err := grpcServer.Serve(listen); err != nil {
		// logger
		return
	}
}
