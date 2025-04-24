package gateway

import (
	"context"
	"github.com/tishntahoe/UrlShortener/internal/storage"
	"github.com/tishntahoe/UrlShortener/proto/redirectpb"
	"github.com/tishntahoe/UrlShortener/proto/shortenerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"
)

type ConnectionGrpcStrct struct {
	ShortenerServiceClient *shortenerpb.ShortenerServiceClient
	RedirectServiceClient  *redirectpb.RedirectServiceClient
}

var Cgs *ConnectionGrpcStrct

func CreateConnectionDial(ipAddress string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		ipAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//logger

		return
	}

}
func GetLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//logger

		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key := r.PathValue("id")
	var st storage.StorageInterface
	link, err := st.Get(ctx, key)
	if err != nil {
		// logger
	}
	w.Write([]byte(link))
	return
}
