package shotener

import pb "github.com/tishntahoe/UrlShortener/proto/shortenerpb"

type Server struct {
	pb.ShortenerServiceServer
}
