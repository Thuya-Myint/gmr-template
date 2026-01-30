package httpserver

import (
	"context"
	"net/http"

	"github.com/Thuya-Myint/gmr-template/internal/middleware"
)

type Server struct {
	httpServer *http.Server
	addr       string
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	var handler http.Handler = mux
	handler = middleware.Logging(handler)
	handler = middleware.Recovery(handler)

	s.httpServer = &http.Server{
		Addr:    s.addr,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
