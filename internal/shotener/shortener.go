package shotener

import (
	"context"
	pb "github.com/tishntahoe/UrlShortener/proto/shortenerpb"
)

type Server struct {
	pb.ShortenerServiceServer
}

func (s Server) ToShort(ctx context.Context, request *pb.ShortRequest) (*pb.ShortResponse, error) {

	return &pb.ShortResponse{}, nil
}
