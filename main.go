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

	mux.HandleFunc("/upload", api.HandleError(a.Upload))
	mux.HandleFunc("/download", api.HandleError(a.Download))
	mux.HandleFunc("/list", api.HandleError(a.List))

	log.Printf("Starting HTTP server at %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
