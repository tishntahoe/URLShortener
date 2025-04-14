package main

import (
	"fmt"
	si "github.com/tishntahoe/UrlShortener/cmd/storageInit"
	gw "github.com/tishntahoe/UrlShortener/internal/gateway"
	cfg "github.com/tishntahoe/UrlShortener/pkg/cfg"
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

	err = gw.CreateConnection(cfgData.ConnectionIpServer)
	if err != nil {
		//logger
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", gw.CreateLinkHandler)
	mux.HandleFunc("/{id}", gw.GetLinkHandler)

	http.ListenAndServe(":8080", mux)
}
