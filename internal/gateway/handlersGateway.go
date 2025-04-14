package gateway

import (
	"context"
	"github.com/tishntahoe/UrlShortener/internal/storage"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type ClientDialConnection struct {
	conn *grpc.ClientConn
}

var DialConn ClientDialConnection

func CreateConnection(ipAddress string) error {
	conn, err := grpc.Dial(ipAddress, grpc.WithTimeout(time.Second*5))
	if err != nil {
		return err
	}
	DialConn = ClientDialConnection{conn}
	return nil
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
