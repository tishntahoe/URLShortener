package main

import (
	"github.com/tishntahoe/UrlShortener/internal/redirect"
	"github.com/tishntahoe/UrlShortener/pkg/logger"
	pb "github.com/tishntahoe/UrlShortener/proto/redirectpb"
	"google.golang.org/grpc"

	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
		return
	}
	grpcServer := grpc.NewServer()

	pb.RegisterRedirectServiceServer(grpcServer, &redirect.Server{})

	if err := grpcServer.Serve(listen); err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
		return
	}
}
