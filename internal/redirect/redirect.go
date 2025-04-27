package redirect

import (
	"context"
	pb "github.com/tishntahoe/UrlShortener/proto/redirectpb"
)

type Server struct {
	pb.RedirectServiceServer
}

func (s Server) ToRedirect(ctx context.Context, request *pb.RedirectShortRequest) (*pb.RedirectShortResponse, error) {

	return &pb.RedirectShortResponse{}, nil
}
