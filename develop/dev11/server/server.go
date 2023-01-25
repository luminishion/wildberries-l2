package server

import (
	"context"
	"log"
	"net/http"

	"dev11/calendar"
	"dev11/server/handlers"
	"dev11/server/middleware"
)

type Server struct {
	sv *http.Server
}

func New(addr string, cl *calendar.Calendar) *Server {
	hl := handlers.NewRouter(cl)
	logMux := middleware.NewLogger(hl.Mux())

	sv := &http.Server{
		Addr:    addr,
		Handler: logMux,
	}

	return &Server{
		sv: sv,
	}
}

func (s *Server) Run() {
	if err := s.sv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func (s *Server) Close(ctx context.Context) {
	if err := s.sv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
