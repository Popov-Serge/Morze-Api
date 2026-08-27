package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	log    *log.Logger
	server *http.Server
}

func GetServer(logger *log.Logger) *Server {
	r := http.NewServeMux()
	r.HandleFunc("/", handlers.Root)
	r.HandleFunc("/upload", handlers.Upload)

	return &Server{
		log: logger,
		server: &http.Server{
			Addr:         ":8080",
			Handler:      r,
			ErrorLog:     logger,
			ReadTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 10,
			IdleTimeout:  time.Second * 15,
		},
	}
}

func (s *Server) ListenAndServe() {
	err := s.server.ListenAndServe()
	if err != nil {
		s.log.Fatal(err)
	}
}
