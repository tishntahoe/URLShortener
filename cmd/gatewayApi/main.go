package main

import (
	"fmt"
	si "github.com/tishntahoe/UrlShortener/cmd/storageInit"
	gw "github.com/tishntahoe/UrlShortener/internal/gateway"
	cfg "github.com/tishntahoe/UrlShortener/pkg/cfg"
	pbRedirect "github.com/tishntahoe/UrlShortener/proto/redirectpb"
	pbShortener "github.com/tishntahoe/UrlShortener/proto/shortenerpb"
	"net/http"
)

func main() {
	cfgData := cfg.CfgLaunch()
	err := si.StorageInit(cfgData)
	if err != nil {
		//logger
		fmt.Println(err)
		return
	}

	shortConn, err := gw.CreateConnectionDial("localhost:50052")    // шортенер cfgData.ConnectionIpServer
	redirectConn, err := gw.CreateConnectionDial("localhost:50051") // редирект cfgData.ConnectionIpServer
	if err != nil {
		//logger
		return
	}
	shortClient := pbShortener.NewShortenerServiceClient(shortConn)
	redirectClient := pbRedirect.NewRedirectServiceClient(redirectConn)

	gw.Cgs = &(gw.ConnectionGrpcStrct{ShortenerServiceClient: &shortClient, RedirectServiceClient: &redirectClient})

	mux := http.NewServeMux()
	mux.HandleFunc("/", gw.CreateLinkHandler)
	mux.HandleFunc("/{id}", gw.GetLinkHandler)
	http.ListenAndServe(":8080", mux)
}
