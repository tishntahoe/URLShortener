package redirect

import (
	"context"
	"github.com/tishntahoe/UrlShortener/internal/storage"
	pb "github.com/tishntahoe/UrlShortener/proto/redirectpb"
)

type Server struct {
	pb.RedirectServiceServer
}

func (s Server) ToRedirect(ctx context.Context, request *pb.RedirectShortRequest) (*pb.RedirectShortResponse, error) {
	rdsconn := *storage.Storage.RedisConn

	Origstr, err := rdsconn.Get(ctx, request.ShortLink).Result()
	if err != nil {
		return nil, err
	}
	rdsconn.Del(ctx, request.ShortLink)
	return &pb.RedirectShortResponse{OrigLink: Origstr}, nil
}
