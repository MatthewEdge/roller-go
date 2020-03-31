package main

import (
	"encoding/json"
	"net/http"
)

type server struct {
	router http.ServeMux
}

func NewServer() *server {
	s := server{}
	s.routes()
	return &s
}

func (s *server) routes() {
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) SendJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
