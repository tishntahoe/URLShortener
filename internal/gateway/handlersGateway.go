package gateway

import (
	"context"
	"github.com/tishntahoe/UrlShortener/internal/storage"
	pb_redirect "github.com/tishntahoe/UrlShortener/proto/redirectpb"
	pb_shortener "github.com/tishntahoe/UrlShortener/proto/shortenerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"strings"
	"time"
)

type ClientDialConnection struct {
	conn *grpc.ClientConn
}

var DialConn *ClientDialConnection

func CreateConnectionDial(ipAddress string) error {
	conn, err := grpc.NewClient(
		ipAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer conn.Close()
	if err != nil {
		return err
	}

	DialConn = &ClientDialConnection{conn}
	return nil
}

func CreateConnectionListener(ipAddress string) error {
	port := strings.Split(ipAddress, ":")[1]

	lis, err := net.Listen("tcp", port)
	if err != nil {
		// logger
		return err
	}

	server := grpc.NewServer()
	pb_redirect.RegisterRedirectServiceServer(server, &pb_redirect.UnimplementedRedirectServiceServer{})
	pb_shortener.RegisterShortenerServiceServer(server, &pb_shortener.UnimplementedShortenerServiceServer{})

	if err := server.Serve(lis); err != nil {
		//logger
		return err
	}

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
