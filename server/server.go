package server

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type server struct {
	router http.ServeMux
}

func NewServer() *server {
	s := &server{}
	s.routes()
	return s
}

func (s *server) routes() {
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.HandleFunc("/roll", s.handleRoll())
}

func (s *server) respond(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
