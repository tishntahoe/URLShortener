package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tishntahoe/UrlShortener/pkg/logger"
	pbRedirect "github.com/tishntahoe/UrlShortener/proto/redirectpb"
	pbShortener "github.com/tishntahoe/UrlShortener/proto/shortenerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"
)

type (
	ConnectionGrpcStrct struct {
		ShortenerServiceClient *pbShortener.ShortenerServiceClient
		RedirectServiceClient  *pbRedirect.RedirectServiceClient
	}
)

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
	if r.Method != "POST" {
		err := errors.New("Неверный метод ")
		logger.InfoHandler(err, logger.GetWorkDir())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var request struct {
		link string `json:"link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
	}
	// обращение к GRPC
	// для получения ссылки СЕТТЕР
	server := *Cgs.ShortenerServiceClient
	resp, err := server.ToShort(ctx, &pbShortener.ShortRequest{
		OrigLink: request.link,
	})
	if err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
	}
	convResp := map[string]string{"short_link": resp.ShortLink}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(convResp)
}

func GetLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("Неверный метод ")
		logger.InfoHandler(err, logger.GetWorkDir())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key := r.PathValue("id")

	// обращение к GRPC
	// для получения ссылки ГЕТТЕР

	server := *Cgs.RedirectServiceClient
	out, err := server.ToRedirect(ctx, &pbRedirect.RedirectShortRequest{ShortLink: key})

	if err != nil {
		logger.ErrorHandler(err, logger.GetWorkDir())
	}

	http.Redirect(w, r, out.OrigLink, http.StatusTemporaryRedirect)
	return
}
