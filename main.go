package main

import (
	"log"
	"net"
	"net/http"

	"file-storage/adapters/storage"
	"file-storage/api"
	"file-storage/config"
)

func main() {
	cfg := config.New()
	mux := http.NewServeMux()
	s := http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler: mux,
	}

	db := storage.New(cfg.Storage)
	a := api.New(cfg.API, db)

	mux.HandleFunc("/files/", api.WrapRouterError(a.Router))
	mux.HandleFunc("/files/list", api.WrapHandlerError(a.List))

	log.Printf("Starting HTTP server at %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
