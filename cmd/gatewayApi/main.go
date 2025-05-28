package main

import (
	"fmt"
	si "github.com/tishntahoe/UrlShortener/cmd/storageInit"
	gw "github.com/tishntahoe/UrlShortener/internal/gateway"
	"github.com/tishntahoe/UrlShortener/internal/redirect"
	"github.com/tishntahoe/UrlShortener/internal/shotener"
	cfg "github.com/tishntahoe/UrlShortener/pkg/cfg"
	pbRedirect "github.com/tishntahoe/UrlShortener/proto/redirectpb"
	pbShortener "github.com/tishntahoe/UrlShortener/proto/shortenerpb"
	"google.golang.org/grpc"
	"net"
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

	go rdrct()
	go shrt()

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

func rdrct() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		// logger
		return
	}
	grpcServer := grpc.NewServer()

	pbRedirect.RegisterRedirectServiceServer(grpcServer, &redirect.Server{})

	if err := grpcServer.Serve(listen); err != nil {
		// logger
		return
	}
}

func shrt() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		// logger
		return
	}
	grpcServer := grpc.NewServer()
	pbShortener.RegisterShortenerServiceServer(grpcServer, &shotener.Server{})

	if err := grpcServer.Serve(listen); err != nil {
		// logger
		return
	}
}
