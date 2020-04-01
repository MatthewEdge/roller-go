package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"roller-go/server"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	rand.Seed(time.Now().UnixNano())

	s := server.NewServer()
	if err := http.ListenAndServe(":8080", s); err != nil {
		return err
	}

	return nil
}
