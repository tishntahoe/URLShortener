package main

import (
	"fmt"
	"github.com/tishntahoe/UrlShortener/cmd/storageInit"
	gw "github.com/tishntahoe/UrlShortener/internal/gateway"
	cfg "github.com/tishntahoe/UrlShortener/pkg/cfg"
	"net/http"
)

func main() {
	cfgData := cfg.CfgLaunch()
	err := storageInit.StorageInit(cfgData)
	if err != nil {
		//logger
		fmt.Println(err)
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", gw.CreateLinkHandler)
	mux.HandleFunc("/{id}", gw.GetLinkHandler)

	http.ListenAndServe(":8080", mux)
}
