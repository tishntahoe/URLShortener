package shotener

import (
	"context"
	"github.com/tishntahoe/UrlShortener/internal/storage"
	"github.com/tishntahoe/UrlShortener/pkg/logger"
	pb "github.com/tishntahoe/UrlShortener/proto/shortenerpb"
	"math/rand"
)

var all string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

type Server struct {
	pb.ShortenerServiceServer
}

func generateRandomLink() string {
	length := 6
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(length, func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

func (s Server) ToShort(ctx context.Context, request *pb.ShortRequest) (*pb.ShortResponse, error) {
	var foundvalue string
	rdsconn := *storage.Storage.RedisConn
	for rdsconn.Get(ctx, foundvalue) != nil || foundvalue == "" {
		foundvalue = generateRandomLink()
	}
	settedVal, err := rdsconn.Set(ctx, foundvalue, request.OrigLink, 5).Result()
	if err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
	}
	return &pb.ShortResponse{ShortLink: settedVal}, nil
}
