package main

import (
	"github.com/tishntahoe/UrlShortener/cmd/storageInit"
	gw "github.com/tishntahoe/UrlShortener/internal/gateway"
	cfg "github.com/tishntahoe/UrlShortener/pkg/cfg"
	"net/http"
)

func main() {
	cfgData := cfg.CfgLaunch()
	db, err := storageInit.StorageInit(cfgData.ConnectionDbString)
	if err != nil {
		// logger
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", gw.CreateLinkHandler)
	mux.HandleFunc("/{id}", gw.GetLinkHandler)

	http.ListenAndServe(":8080", mux)
}
