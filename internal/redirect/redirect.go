package redirect

import pb "github.com/tishntahoe/UrlShortener/proto/redirectpb"

type Server struct {
	pb.RedirectServiceServer
}
